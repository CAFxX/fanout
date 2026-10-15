package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16 }
func genFn0() string {
    v6 := []int{32767, 2147483647, 0, 32768}
    v7 := GenS1{F0: 1, F1: ("foo bar" + ("hello" + (("" + ((("a" + "hello") + ("\n\t" + "foo bar")) + "\n\t")) + ""))), F2: 128}
    var v8 int64 = 9223372036854775807
    v9 := GenS2{F0: 2, F1: 702}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (("" + "hello") + "a")
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 int8 = -1
    var v5 bool = false
    var v10 float32 = 0.1
    var v11 int8 = (v4 ^ -128)
    var v12 int8 = v4
    v11 = 127
    fmt.Println(v4, v11)
    var v13 int8 = (v4 + v12)
    var v14 uint16 = 42
    v14 -= 255
    switch "\n\t" {
    case "foo bar":
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 uint32 = 4294967295
            _ = fi15
            _ = v16
        }
    }
    switch 42 {
    case 2147483647:
        v17 := GenS1{F0: 593, F1: "\n\t", F2: 255}
        _ = v17
    case -662:
        switch ("a" + "foo bar") {
        case "foo bar":
            for fi18 := 0; fi18 < 7; fi18++ {
                if v5 {
                    switch -2147483648 {
                    case -929:
                        v13 += (v4 | v4)
                    case -1:
                        if false {
                            var v19 float32 = (float32((-2.25 / (0.1 - 1.5))) * (v10 - float32(float64(((((float32(1.5) + v10) * v10) - v10) - 0.1)))))
                            _ = v19
                        }
                    default:
                        v20 := GenS3{F0: 1000}
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v22 := []int{}
                        _ = fi21
                        _ = v22
                    }
                    v23 := GenS2{F0: 65535, F1: 4294967295}
                    _ = v23
                }
                v14 = (^(v14))
                v24 := v3["\n\t"]
                _ = fi18
                _ = v24
            }
        }
    case -205:
        var v25 float32 = 1e10
        _ = v25
    }
    v5 = true
    v10 = ((v10 + (0.0 - 0.0)) * float32((3.14159 - 0.1)))
    var v26 float32 = v10
    for fi27 := 0; fi27 < 18; fi27++ {
        fmt.Println(v11, v13, v5)
        v4 += (v11 / ((42) | 1))
        _ = fi27
    }
    fmt.Println(v10, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
