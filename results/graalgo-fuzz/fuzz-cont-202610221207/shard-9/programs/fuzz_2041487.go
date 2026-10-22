package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float64 = (float64((((float32(((3.14159 + (float64(3.14159) * ((1e10 - (1.5 * 3.14159)) * 1e10))) * 3.14159)) - 0.1) + 1.5) + (-2.25 / ((0.0 + 0.1) * (((1e10 * 1.5) - -2.25) - (3.14159 * 0.0)))))) * (float64(3.14159) + -2.25))
    var v2 int = 2147483647
    v3 := make(map[uint32]int)
    var v4 uint16 = 2
    v5 := make(map[string]string)
    var v6 uint64 = 1
    switch ("\n\t" + (("foo bar" + ((("a" + "\n\t") + "foo bar") + (("" + "foo bar") + ("hello" + "a")))) + "")) {
    case "foo bar":
        delete(v5, "a")
    }
    if (!false) {
        v7 := []int{}
        var v8 uint = 266
        _ = v7
        _ = v8
    }
    if ((-2.25 / ((1.5 / 1.0) * 3.14159)) <= (0.1 - 0.0)) {
        if true {
            v4 = 65535
        }
        v1 = v1
    }
    var v9 int64 = -1
    v4 = 676
    if false {
        if false {
            switch ((("foo bar" + "foo bar") + "a") + "\n\t") {
            case "hello":
                v6 += 0
            case "\n\t":
                v1 -= v1
            default:
                if false {
                    v10 := []int{(^(v2))}
                    for fi11 := 0; fi11 < 5; fi11++ {
                        fmt.Println(fi11)
                        if (!(0 > 1)) {
                            for fi12 := 0; fi12 < 3; fi12++ {
                                v10 = append(v10, -9223372036854775808)
                                var v13 uint64 = 0
                                _ = fi12
                                _ = v13
                            }
                            var v14 int64 = 97
                            _ = v14
                        } else {
                            v6 += 1
                        }
                        v15 := GenS1{F0: 127}
                        _ = fi11
                        _ = v15
                    }
                    _ = v10
                }
            }
        }
        var v16 uint = 42
        _ = v16
    }
    fmt.Println(v2)
    var v17 int8 = 42
    fmt.Println(v4, v6)
    var v18 float64 = v1
    v1 = v1
    v19 := GenS1{F0: (v17 / ((42) | 1))}
    v17 = v17
    var v20 uint64 = 350
    v20 += (v6 + 42)
    v21 := []int{v2, (v2 / (((v2 & (^(v2)))) | 1)), v2, -2147483648}
    fmt.Println(v2, v4)
    v22 := GenS1{F0: v17}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
