package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 int32 = 0
    v4 := make(map[int]string)
    var v5 int = -1
    var v6 int32 = v3
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            v8 := GenS1{F0: ""}
            _ = v8
        }
        v1["foo bar"] = v5
        if (!(1 == 4294967295)) {
            if ((v3 >> ((0) & 63)) > v6) {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 float32 = 1.5
                    _ = fi9
                    _ = v10
                }
            }
        }
        _ = fi7
    }
    fmt.Println(v3, v5)
    v3 = (v6 * v6)
    switch "hello" {
    case "foo bar":
        var v11 float32 = 3.14159
        _ = v11
    case "a":
        var v12 uint = 42
        _ = v12
    default:
        var v13 float32 = (3.14159 * 1.5)
        _ = v13
    }
    fmt.Println(v6, v6, v3)
    v14 := []int{}
    for fi15 := 0; fi15 < 17; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            switch fi15 {
            case 1:
                fmt.Println(fi16, v6)
            }
            _ = fi16
        }
        v5 -= 32768
        v4[fi15] = ("\n\t" + "a")
        _ = fi15
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS2{F0: 42}
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 859
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
