package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int16) int {
    if ((2147483648 <= 2147483648) && (4294967295 <= 361)) {
        var v6 float64 = float64(float32((-2.25 - 1e10)))
        fmt.Println(p0)
        v7 := GenS1{F0: -9223372036854775808}
        _ = v6
        _ = v7
    }
    v8 := []int{}
    fmt.Println(p0)
    v9 := GenS2{F0: 2}
    _ = p0
    _ = v8
    _ = v9
    return 32767
}
func genFn1(p0 string) int64 {
    p0 = ""
    var v10 string = "a"
    for fi11 := 0; fi11 < 17; fi11++ {
        if (65535 != 2147483648) {
            for fi12 := 0; fi12 < 5; fi12++ {
                var v13 int8 = 127
                switch v10 {
                case "foo bar":
                    v14 := []int{(^(fi11)), -1, fi12, 1}
                    _ = v14
                }
                p0 = v10
                _ = fi12
                _ = v13
            }
        }
        _ = fi11
    }
    var v15 uint = 18446744073709551615
    _ = p0
    _ = v10
    _ = v15
    return -1
}

func main() {
    var v1 string = (("" + "\n\t") + ((("hello" + ("" + "a")) + "") + ""))
    var v2 uint16 = 1
    var v3 bool = false
    var v4 uint = 597
    var v5 float64 = (3.14159 - 1.5)
    var v16 bool = true
    if true {
        fmt.Println(v1, v4, v4)
    } else {
        v3 = (v1 == (("" + "a") + (v1 + v1)))
        v17 := []int{32767, 2147483647, -9223372036854775808}
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint = 432
        _ = fi18
        _ = v19
    }
    v20 := []int{1, 0, 0, -883}
    fmt.Println(v16)
    fmt.Println(v3, v4)
    if (v5 < 3.14159) {
        v20 = append(v20, 4294967295)
        if len(v20) > 0 {
            v20[0] = genFn0(-32768)
        }
    }
    _ = len(v20)
    var v21 bool = (2 <= 18446744073709551615)
    var v22 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
