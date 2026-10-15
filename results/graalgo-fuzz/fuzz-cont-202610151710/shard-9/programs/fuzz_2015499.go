package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64 }
func genFn0(p0 int, p1 int32, p2 string) int8 {
    v4 := []int{(^(p0))}
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -40
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    v3 := make(map[int]uint32)
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 float32 = ((-2.25 - (float32(0.0) * (float32(1e10) / float32(1e10)))) / (0.1 + 1.5))
        v7 := GenS1{F0: 42, F1: ""}
        if (!(42 < 1)) {
            v6 -= -2.25
            v1[""] = 162
            switch "a" {
            case "hello":
                var v8 float32 = (0.0 * ((v6 - -2.25) - float32(1.5)))
                _ = v8
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi5, fi5, v6)
                default:
                    switch fi5 {
                    case 32767:
                        var v9 float64 = float64((((-2.25 - v6) * v6) + (0.0 + (((3.14159 + (0.1 / v6)) / (v6 - 1e10)) - (1.5 * (float32(float64(v6)) - v6))))))
                        _ = v9
                    case 609:
                        v2 = v2
                    default:
                        var v10 int64 = 9223372036854775807
                        _ = v10
                    }
                }
            }
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 float32 = v6
                var v13 uint = 242
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    fmt.Println(v2)
    var v14 uint8 = 127
    fmt.Println(v14, v2)
    for fi15 := 0; fi15 < 5; fi15++ {
        if (float32(float64(((float32(1.5) + 1e10) - (1.5 / ((-2.25 + 1.5) / 3.14159))))) > 3.14159) {
            switch -204 {
            case -318:
                v16 := []int{fi15, fi15, (fi15 + 505)}
                _ = v16
            case -1:
                var v17 string = ""
                _ = v17
            }
            if false {
                v18 := GenS2{F0: 730}
                _ = v18
            }
            v2 = -45
        } else {
            fmt.Println(v14, v14, fi15)
            if false {
                v19 := GenS1{F0: 681, F1: "foo bar"}
                _ = v19
            }
        }
        delete(v1, "")
        _ = fi15
    }
    var v20 float64 = 0.1
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(v14)
        fmt.Println(v14, v20)
        v22 := GenS2{F0: 9223372036854775807}
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: -9223372036854775808}
    for fi24 := 0; fi24 < 16; fi24++ {
        v20 = v20
        for fi25 := 0; fi25 < 2; fi25++ {
            v14 -= v14
            _ = fi25
        }
        v26 := []int{(fi24 | 0), -307, fi24}
        _ = fi24
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
