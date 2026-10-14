package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() int {
    v7 := GenS1{F0: 161}
    var v8 int64 = 42
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int8 = 60
        fmt.Println(v10, fi9, v10)
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 1}
    v12 := GenS1{F0: 2147483648}
    _ = v11
    _ = v12
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 42
    var v3 int8 = v2
    var v4 float32 = 3.14159
    v5 := make(map[string]bool)
    v6 := []int{2147483647, -897, 32767}
    var v13 float32 = (v4 * (v4 * v4))
    v14 := GenS1{F0: 609}
    fmt.Println(v2)
    v15 := GenS1{F0: 65535}
    fmt.Println(v4, v3)
    var v16 int = -1
    var v17 uint64 = 1
    var v18 float64 = (0.1 * 1e10)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        if false {
            fmt.Println(v16, v3, v18)
            if true {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v16 += fi19
            } else {
                if (v16 <= (-(fi19))) {
                    switch fi19 {
                    case -685:
                        v18 = (float64(float32((3.14159 + (v18 / (3.14159 * v18))))) * (0.1 + (v18 / (v18 - (v18 - float64(1e10))))))
                    default:
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v1[2147483647] = 699
                            _ = fi20
                        }
                    }
                } else {
                    v5["\n\t"] = false
                    fmt.Println(v18, v3, fi19)
                }
            }
            v13 = 3.14159
        }
        switch (("" + (("foo bar" + "foo bar") + "\n\t")) + "\n\t") {
        case "foo bar":
            v13 = 1e10
        case "":
            v21 := v1[(v16 | 127)]
            _ = v21
        case "hello":
            _ = len(v5)
        }
        var v22 uint64 = 2
        _ = fi19
        _ = v22
    }
    v23 := GenS1{F0: 2147483647}
    delete(v5, "")
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 int = (-(v16))
        _ = fi24
        _ = v25
    }
    if true {
        _ = len(v6)
    }
    var v26 int = genFn0()
    v27 := GenS1{F0: uint32(v17)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
