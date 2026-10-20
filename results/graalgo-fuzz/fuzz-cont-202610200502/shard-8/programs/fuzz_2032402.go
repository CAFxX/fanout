package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 bool) uint {
    fmt.Println(p0)
    if true {
        if (p0 && ((!(9223372036854775807 >= -9223372036854775808)) || false)) {
            if p0 {
                if true {
                    var v7 int32 = 975
                    _ = v7
                } else {
                    if true {
                        v8 := GenS1{F0: 125}
                        v9 := []int{32768, -1, 405}
                        var v10 float64 = (float64(-2.25) / float64(-2.25))
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                    fmt.Println(p0)
                }
            }
            v11 := GenS2{F0: 33}
            _ = v11
        } else {
            var v12 int32 = -2147483648
            _ = v12
        }
        switch "a" {
        case "hello":
            if (2147483648 >= 349) {
                var v13 string = (("" + ("a" + "\n\t")) + "foo bar")
                _ = v13
            } else {
                v14 := []int{42, 9223372036854775807}
                _ = v14
            }
        case "foo bar":
            switch "" {
            case "":
                var v15 float64 = -2.25
                _ = v15
            case "\n\t":
                v16 := []int{127, 844}
                _ = v16
            case "hello":
                var v17 int64 = 1
                _ = v17
            }
        }
        fmt.Println(p0)
    } else {
        if (737 > 9223372036854775807) {
            fmt.Println(p0)
            v18 := GenS2{F0: -206}
            _ = v18
        }
    }
    v19 := []int{}
    for fi20 := 0; fi20 < 1; fi20++ {
        p0 = p0
        _ = fi20
    }
    _ = p0
    _ = v19
    return 970
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 1
    var v3 int64 = 42
    var v4 float64 = ((((0.1 - (0.1 - 0.1)) - 1e10) - ((float64(1e10) - ((1e10 * (1e10 + float64(-2.25))) - 1.5)) / float64(3.14159))) + 3.14159)
    var v5 int8 = 1
    var v6 uint16 = v2
    if ((v3 + v3) < v3) {
        var v21 int64 = -9223372036854775808
        var v22 uint8 = 255
        _ = v21
        _ = v22
    }
    v5 -= 102
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 int = int(v3)
        _ = fi23
        _ = v24
    }
    var v25 uint64 = 2
    if true {
        v26 := []int{2147483647, -906}
        v27 := v1[65535]
        _ = v26
        _ = v27
    } else {
        var v28 uint = 65535
        fmt.Println(v5, v3, v6)
        if true {
            var v29 int16 = 399
            v25 = 555
            _ = v29
        }
        _ = v28
    }
    var v30 float64 = v4
    v31 := GenS2{F0: 0}
    v32 := GenS2{F0: 730}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
