package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 string, p1 uint32) int16 {
    if (false && true) {
        var v6 bool = (!(42 < -114))
        _ = v6
    }
    var v7 string = "foo bar"
    _ = p0
    _ = p1
    _ = v7
    return -554
}

func main() {
    var v1 string = "foo bar"
    var v2 string = v1
    var v3 int32 = 1
    v4 := []int{127}
    var v5 int64 = 1
    v5 = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v8 float32 = (-2.25 + -2.25)
    var v9 uint64 = 731
    if (!false) {
        for fi10 := 0; fi10 < 19; fi10++ {
            v11 := GenS1{F0: 0}
            fmt.Println(v9)
            _ = fi10
            _ = v11
        }
    }
    if ((v8 + (((1.5 / (float32((0.1 / float64(float32((-2.25 * 1.5))))) - v8)) / v8) - 3.14159)) >= 1e10) {
        switch -520 {
        case -836:
            fmt.Println(v8, v5, v2)
        case 238:
            v12 := GenS1{F0: v5}
            _ = v12
        default:
            for fi13 := 0; fi13 < 11; fi13++ {
                var v14 uint16 = 0
                var v15 float64 = 0.0
                if (2147483647 == (v3 & v3)) {
                    v16 := GenS1{F0: (v5 * int64(fi13))}
                    v17 := GenS1{F0: v5}
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v14
                _ = v15
            }
        }
        switch ((("foo bar" + "hello") + (v1 + ((v2 + "foo bar") + (v1 + v2)))) + "hello") {
        case "a":
            if (465 <= 2) {
                fmt.Println(v8)
                for fi18 := 0; fi18 < 14; fi18++ {
                    v1 = ""
                    _ = len(v4)
                    _ = fi18
                }
                fmt.Println(v8, v3, v8)
            }
        case "\n\t":
            var v19 uint64 = 797
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 float64 = 1e10
            v1 = v2
            _ = fi20
            _ = v21
        }
        var v22 int32 = v3
        var v23 uint16 = 0
        _ = v22
        _ = v23
    }
    switch "hello" {
    case "a":
        switch v1 {
        case "":
            fmt.Println(v9, v3, v5)
        case "a":
            switch "a" {
            case "a":
                var v24 int64 = 9223372036854775807
                _ = v24
            default:
                fmt.Println(v1, v2)
            }
        case "foo bar":
            var v25 uint32 = 681
            _ = v25
        }
    case "foo bar":
        if true {
            var v26 uint64 = 35
            fmt.Println(v26, v5, v9)
            _ = v26
        } else {
            var v27 uint16 = 510
            switch 309 {
            case -2147483648:
                v3 += (-(v3))
            }
            _ = v27
        }
    case "\n\t":
        var v28 uint = 0
        _ = v28
    }
    if ((v3 * (v3 ^ 1)) > (-(v3))) {
        v29 := GenS1{F0: v5}
        for fi30 := 0; fi30 < 8; fi30++ {
            fmt.Println(v5, v1, v3)
            v4 = append(v4, fi30)
            _ = fi30
        }
        _ = v29
    }
    if (false || true) {
        if (!(1 >= 2)) {
            var v31 uint64 = v9
            _ = v31
        }
        v1 = ((v1 + v1) + "a")
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
