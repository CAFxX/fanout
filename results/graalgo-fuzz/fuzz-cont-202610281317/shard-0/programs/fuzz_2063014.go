package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 2
    var v3 uint16 = 939
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    v1 += v1
    v1 += v1
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := []int{fi6}
        var v8 float32 = (float32(((3.14159 * 1.5) + ((((0.1 - 0.1) - 1e10) / 3.14159) * ((((-2.25 / 1e10) + 3.14159) - (3.14159 * 3.14159)) - (3.14159 - 3.14159))))) + float32(float64((((1e10 - (3.14159 + float32(float64(0.0)))) * -2.25) + float32((3.14159 / (0.1 * 1.5)))))))
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        if false {
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := []int{(fi9 ^ (fi10 + -954)), -1, fi10, 127}
                _ = fi10
                _ = v11
            }
            switch ("foo bar" + "hello") {
            case "a":
                var v12 float64 = ((3.14159 / (1.5 / 0.1)) + -2.25)
                _ = v12
            }
        }
        _ = fi9
    }
    v1 -= uint8(v2)
    switch "\n\t" {
    case "\n\t":
        var v13 int16 = -1
        _ = v13
    case "foo bar":
        v1 -= (v1 - (^(v1)))
    default:
        v1 += 2
    }
    if (1 < 42) {
        if true {
            if false {
                var v14 uint = 0
                for fi15 := 0; fi15 < 4; fi15++ {
                    for fi16 := 0; fi16 < 8; fi16++ {
                        if true {
                            v1 += 28
                        } else {
                            v1 = v1
                        }
                        v17 := []int{fi15, fi15, fi15}
                        v5["a"] = ("hello" + "")
                        _ = fi16
                        _ = v17
                    }
                    v18 := []int{(fi15 | fi15), (fi15 + fi15), (fi15 & fi15)}
                    _ = fi15
                    _ = v18
                }
                var v19 string = ""
                _ = v14
                _ = v19
            }
        }
        var v20 uint32 = 212
        v21 := GenS1{F0: (^(v1))}
        _ = v20
        _ = v21
    } else {
        switch "\n\t" {
        case "foo bar":
            var v22 uint16 = v3
            _ = v22
        case "hello":
            if (!true) {
                v23 := GenS1{F0: v1}
                var v24 uint8 = 112
                var v25 float64 = 1.5
                _ = v23
                _ = v24
                _ = v25
            }
        }
    }
    v26 := GenS1{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
