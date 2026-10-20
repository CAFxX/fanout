package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64; F1 string; F2 uint32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int64) bool {
    v4 := []int{803, -763, 128}
    var v5 string = "\n\t"
    fmt.Println(p0)
    _ = p0
    _ = v4
    _ = v5
    return (p0 != (p0 * 1))
}
func genFn1(p0 uint, p1 string, p2 int64) uint {
    fmt.Println(p1, p2, p2)
    if (182 != 38) {
        fmt.Println(p0, p1)
        v6 := []int{-1, 127}
        if ((-1 > 32767) || (!true)) {
            p0 += p0
            var v7 int64 = -1
            switch (("a" + (p1 + p1)) + p1) {
            case "\n\t":
                p0 -= p0
            }
            _ = v7
        } else {
            v8 := GenS1{F0: -596}
            _ = v8
        }
        _ = v6
    }
    p1 = p1
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    var v1 int64 = 764
    var v2 string = ("\n\t" + (("foo bar" + ("hello" + "foo bar")) + "a"))
    v3 := []int{0, -9223372036854775808, 42}
    var v10 float64 = (0.0 - (3.14159 * ((float64(1.5) / 1.0) - -2.25)))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(fi11)
        _ = fi11
    }
    switch "foo bar" {
    case "":
        var v12 int16 = -1
        _ = v12
    default:
        v10 += (-2.25 - float64((3.14159 + 1.5)))
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v3 = append(v3, fi13)
        var v14 bool = true
        _ = fi13
        _ = v14
    }
    switch int(v1) {
    case 127:
        v1 = 656
    }
    v2 = v2
    if false {
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                switch (^(fi16)) {
                case 127:
                    var v17 int8 = -5
                    _ = v17
                default:
                    var v18 float32 = 1e10
                    v18 = ((v18 - v18) * v18)
                    if false {
                        _ = len(v3)
                    } else {
                        v18 += (0.0 * (v18 - 0.1))
                        v2 = v2
                        var v19 uint32 = 775
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v20 := GenS1{F0: -840}
                        v18 -= 0.1
                        _ = v19
                        _ = v20
                    }
                    _ = v18
                }
                if true {
                    var v21 uint32 = 2147483648
                    _ = v21
                }
                _ = fi16
            }
            _ = fi15
        }
        switch -9223372036854775808 {
        case 9223372036854775807:
            v22 := GenS2{F0: -160, F1: 769, F2: 127}
            _ = v22
        default:
            switch (((v2 + (v2 + "a")) + (v2 + v2)) + ("foo bar" + (((((v2 + v2) + v2) + v2) + (v2 + "hello")) + "hello"))) {
            case "a":
                var v23 float32 = 1e10
                _ = v23
            }
        }
        v2 = ("hello" + v2)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v24 := GenS2{F0: 32767, F1: 18446744073709551615, F2: 127}
    switch 32768 {
    case 120:
        v3 = append(v3, 404)
    case 198:
        if (!((v2 > v2) || ((42 >= 1) && (false && (0 != 18446744073709551615))))) {
            fmt.Println(v1)
            v25 := GenS3{F0: 2, F1: "", F2: 2}
            v2 = ("hello" + "a")
            _ = v25
        }
    case 1:
        v26 := GenS4{F0: 2147483647}
        _ = v26
    default:
        switch -2147483648 {
        case 4294967295:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        for fi28 := 0; fi28 < 15; fi28++ {
            v29 := GenS4{F0: 0}
            var v30 float32 = 3.14159
            _ = fi28
            _ = v29
            _ = v30
        }
        var v31 int32 = -1
        var v32 uint8 = 42
        _ = fi27
        _ = v31
        _ = v32
    }
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
}
