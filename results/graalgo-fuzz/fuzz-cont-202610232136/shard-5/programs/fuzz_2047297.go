package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []string{"\n\t", "\n\t", "a"}
    var v2 int = 9223372036854775807
    var v3 uint = 255
    v4 := GenS1{F0: 1}
    for fi5 := 0; fi5 < 1; fi5++ {
        v1 = append(v1, ("\n\t" + "a"))
        _ = fi5
    }
    v6 := GenS1{F0: 2}
    if (true && false) {
        var v7 int8 = 100
        _ = v7
    }
    v3 -= v3
    v8 := GenS1{F0: 0}
    v3 += (v3 ^ v3)
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            _ = len(v1)
            if (v3 == v3) {
                fmt.Println(fi9, v2, fi9)
                var v11 float64 = (0.1 - -2.25)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v3 += v3
                _ = v11
            } else {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
            }
            if true {
                fmt.Println(v2, v2)
            }
            _ = fi10
        }
        v2 = (v2 + v2)
        v3 = v3
        _ = fi9
    }
    v2 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
