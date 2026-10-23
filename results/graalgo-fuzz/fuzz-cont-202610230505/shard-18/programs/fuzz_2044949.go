package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0() bool {
    var v7 uint32 = 676
    fmt.Println(v7)
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 uint32 = 2147483647
        var v10 uint32 = 1
        if true {
            for fi11 := 0; fi11 < 11; fi11++ {
                v9 = 0
                if (float64(-2.25) >= ((1.5 * 0.0) + (float64(0.0) + 1e10))) {
                    var v12 int64 = 42
                    var v13 int64 = (-(v12))
                    v12 = (v12 ^ v12)
                    _ = v12
                    _ = v13
                }
                var v14 bool = (-1 >= 745)
                _ = fi11
                _ = v14
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    if (9223372036854775807 > 1) {
        v15 := []int{32767}
        _ = v15
    } else {
        var v16 uint16 = 0
        _ = v16
    }
    _ = v7
    return (("a" + "hello") == "")
}
func genFn1(p0 int8, p1 bool, p2 bool) uint16 {
    if false {
        if p1 {
            for fi17 := 0; fi17 < 18; fi17++ {
                p1 = true
                var v18 float64 = 0.0
                v19 := GenS1{F0: 2}
                _ = fi17
                _ = v18
                _ = v19
            }
        } else {
            p1 = true
            fmt.Println(p1)
            p0 = (p0 << ((4294967295) & 63))
        }
    }
    if ("a" <= "a") {
        var v20 int32 = 0
        var v21 uint8 = 1
        _ = v20
        _ = v21
    } else {
        var v22 string = "a"
        var v23 uint64 = 18446744073709551615
        if ((("foo bar" + "a") + v22) != "foo bar") {
            var v24 uint64 = (v23 ^ v23)
            _ = v24
        }
        _ = v22
        _ = v23
    }
    var v25 int8 = p0
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 int = fi26
        fmt.Println(v27, p0, fi26)
        _ = fi26
        _ = v27
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 63
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]string)
    v3 := []string{"foo bar", ("foo bar" + (("foo bar" + "a") + "a"))}
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    v6 := []string{}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if ((2 >= 42) && ((genFn1(0, false, false) > 303) && ((!((false || true) && false)) || false))) {
        for fi28 := 0; fi28 < 13; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                if false {
                    var v30 uint16 = genFn1(105, (false || true), ((!(-9223372036854775808 >= fi29)) && false))
                    v30 -= genFn1(124, true, (0.0 == float64((0.1 * (float32(float64(3.14159)) / 1.0)))))
                    v30 = 255
                    fmt.Println(fi28, v30, fi28)
                    _ = v30
                }
                v3 = append(v3, ("a" + "a"))
                var v31 uint16 = 567
                _ = fi29
                _ = v31
            }
            _ = fi28
        }
        var v32 uint32 = 65535
        _ = v32
    }
    for fi33 := 0; fi33 < 3; fi33++ {
        if true {
            v3 = append(v3, ("" + ("hello" + "a")))
            delete(v1, 65535)
            if (0 < genFn1(-128, (true && false), true)) {
                var v34 int16 = 802
                v34 = v34
                _ = v34
            }
        }
        _ = fi33
    }
    var v35 int32 = 1
    fmt.Println(v35)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v5, 231)
    v35 -= (v35 / ((v35) | 1))
    fmt.Println(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
