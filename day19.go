package main

import (
	"regexp"
	"strconv"
	"strings"
)

type ruleType int8

const (
	greaterThen = iota
	lessThen
	catch
)

type workflow struct {
	rules []rule
}

type rule struct {
	t    ruleType
	part string
	num  int
	goal string
}

func mainDay19(in string) string {
	inputParts := strings.Split(in, "\n\n")
	works := strings.Split(inputParts[0], "\n")

	workflows := make(map[string]workflow)

	var re = regexp.MustCompile(`(?m)(.*){(.*)}`)

	for _, work := range works {
		matches := re.FindStringSubmatch(work)

		name := matches[1]
		rules := strings.Split(matches[2], ",")

		workflowRules := make([]rule, 0)
		for _, r := range rules {
			rSplit := strings.Split(r, ":")
			if len(rSplit) <= 1 {
				workflowRules = append(workflowRules, rule{
					t:    catch,
					goal: rSplit[0],
				})
				continue
			}
			lt := strings.Split(rSplit[0], "<")
			if len(lt) > 1 {
				n, _ := strconv.Atoi(lt[1])

				workflowRules = append(workflowRules, rule{
					t:    lessThen,
					part: lt[0],
					num:  n,
					goal: rSplit[1],
				})
				continue
			}
			gt := strings.Split(rSplit[0], ">")
			if len(gt) > 1 {
				n, _ := strconv.Atoi(gt[1])

				workflowRules = append(workflowRules, rule{
					t:    greaterThen,
					part: gt[0],
					num:  n,
					goal: rSplit[1],
				})
				continue
			}
			panic("unable to parse rule: " + r)
		}

		workflows[name] = workflow{workflowRules}
	}

	//fmt.Println(workflows)

	partStrings := strings.Split(inputParts[1], "\n")
	parts := make([]map[string]int, 0)

	for _, partString := range partStrings {
		attrs := map[string]int{}

		for _, attr := range strings.Split(partString[1:len(partString)-1], ",") {
			name := attr[0:1]

			num, _ := strconv.Atoi(attr[2:])

			attrs[name] = num
		}

		parts = append(parts, attrs)
	}

	sum := 0

	for _, part := range parts {
		currentWorkflowName := "in"

	runWorkflow:
		for {
			if currentWorkflowName == "R" {
				break
			}
			if currentWorkflowName == "A" {
				sum += part["x"] + part["m"] + part["a"] + part["s"]
				break
			}
			for _, r := range workflows[currentWorkflowName].rules {
				switch r.t {
				case lessThen:
					if part[r.part] < r.num {
						currentWorkflowName = r.goal
						continue runWorkflow
					}
				case greaterThen:
					if part[r.part] > r.num {
						currentWorkflowName = r.goal
						continue runWorkflow
					}
				case catch:
					currentWorkflowName = r.goal
					continue runWorkflow
				}
			}
		}
	}

	return strconv.Itoa(sum)
}
