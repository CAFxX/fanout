package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 uint8, p1 int) bool {
    fmt.Println(p1)
    switch "\n\t" {
    case "foo bar":
        for fi4 := 0; fi4 < 17; fi4++ {
            p0 -= p0
            _ = fi4
        }
    case "a":
        p1 -= (p1 % ((p1) | 1))
    }
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int8 = 1
        _ = fi5
        _ = v6
    }
    v7 := []int{127}
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 int = 44
        _ = fi8
        _ = v9
    }
    _ = len(v2)
    var v10 uint64 = 332
    if false {
        v2[(("\n\t" + "a") + "")] = false
        var v11 int16 = -1
        _ = v11
    }
    fmt.Println(v10)
    v12 := GenS1{F0: (!false)}
    v10 = v10
    var v13 string = ""
    var v14 int = 32768
    v13 = v13
    if true {
        if (true || (true && true)) {
            v13 = ((v13 + (v13 + v13)) + "a")
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
