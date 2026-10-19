package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS5 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 uint64) int64 {
    v7 := GenS1{F0: 2, F1: ("foo bar" + "foo bar")}
    if false {
        switch 9223372036854775807 {
        case -1:
            var v8 int8 = 0
            _ = v8
        case 329:
            for fi9 := 0; fi9 < 16; fi9++ {
                p0 = p0
                v10 := []int{(fi9 & fi9)}
                _ = fi9
                _ = v10
            }
        }
    } else {
        v11 := GenS1{F0: 1, F1: "hello"}
        var v12 uint16 = 2
        switch ("\n\t" + ("" + (("foo bar" + "hello") + ("a" + "")))) {
        case "foo bar":
            var v13 float64 = (-2.25 - 0.1)
            _ = v13
        case "":
            v14 := []int{2147483647, 293, -953}
            _ = v14
        default:
            var v15 int8 = 1
            _ = v15
        }
        _ = v11
        _ = v12
    }
    p2 -= (p2 >> ((958) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 bool, p1 int) int32 {
    v16 := GenS2{F0: 65535}
    var v17 int64 = 0
    var v18 uint64 = 42
    switch (("foo bar" + "hello") + (("hello" + ("" + "foo bar")) + ("foo bar" + "\n\t"))) {
    case "hello":
        v19 := []int{}
        _ = v19
    case "\n\t":
        if ((false && p0) || p0) {
            v20 := []int{40}
            _ = v20
        }
    case "a":
        v21 := []int{int(v18), p1}
        _ = v21
    default:
        v22 := GenS3{F0: "\n\t", F1: 423, F2: 65535}
        _ = v22
    }
    v23 := []int{p1, p1, p1, p1}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    _ = v23
    return 425
}

func main() {
    var v1 int = 32767
    var v2 int32 = 1
    var v3 uint = 65535
    var v4 uint = (v3 >> ((v3) & 63))
    v5 := []string{"hello", "foo bar", ("" + ("a" + (("\n\t" + ("" + "hello")) + ("foo bar" + (("hello" + ("a" + "hello")) + "a")))))}
    var v6 string = ("" + "foo bar")
    var v24 uint = (v3 + (v3 - 1))
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS4{F0: 2, F1: 65535, F2: -68}
        if (42 > 42) {
            for fi27 := 0; fi27 < 18; fi27++ {
                v1 = v1
                var v28 string = v6
                _ = fi27
                _ = v28
            }
            var v29 float64 = 0.1
            v29 += v29
            _ = v29
        } else {
            var v30 uint = v3
            fmt.Println(v3)
            _ = v30
        }
        v24 -= 18446744073709551615
        _ = fi25
        _ = v26
    }
    if len(v5) > 0 {
        v5[0] = v6
    }
    if false {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        fmt.Println(v1, v3)
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v24 = v4
    var v31 float64 = float64(0.0)
    v6 = v6
    if false {
        if (!true) {
            v32 := GenS5{F0: v2, F1: 2147483648}
            for fi33 := 0; fi33 < 15; fi33++ {
                if len(v5) > 0 {
                    v5[0] = (v6 + "\n\t")
                }
                var v34 bool = true
                if ((127 <= 2) || ((true || (32767 < 1)) || (!("foo bar" <= v6)))) {
                    v35 := GenS3{F0: v6, F1: 167, F2: 2}
                    _ = v35
                }
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        v6 = "foo bar"
        for fi36 := 0; fi36 < 19; fi36++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi36
        }
    } else {
        var v37 float32 = ((3.14159 + 3.14159) / (0.1 - 1e10))
        v2 = genFn1((65535 >= 0), (v1 << ((v3) & 63)))
        _ = v37
    }
    fmt.Println(v24, v31)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
