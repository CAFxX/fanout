package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint32, p1 uint64, p2 int32) string {
    v6 := []int{128, 127, -647}
    v7 := GenS1{F0: 1}
    var v8 uint = 158
    p1 += uint64(v8)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return ("foo bar" + "foo bar")
}
func genFn1() uint8 {
    v9 := []int{39, 9223372036854775807}
    switch 34 {
    case 32768:
        var v10 uint16 = 0
        _ = v10
    case 350:
        v11 := GenS1{F0: 32768}
        _ = v11
    }
    _ = v9
    return 19
}

func main() {
    v1 := []string{"hello", "a"}
    var v2 bool = true
    var v3 bool = (65535 >= 535)
    v4 := []int{411}
    var v5 uint16 = 823
    if v3 {
        var v12 int32 = 1
        v5 = v5
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 bool = v3
        var v15 int16 = 0
        if ((^(v5)) <= 2) {
            if true {
                v1 = append(v1, (("hello" + (("\n\t" + "a") + ("" + "foo bar"))) + ("foo bar" + "\n\t")))
                v15 = v15
                v14 = true
            }
            var v16 int32 = -177
            _ = v16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        _ = fi17
    }
    v2 = (v2 || (0 > 1000))
    if (459 == 0) {
        var v18 int64 = 42
        var v19 uint32 = 0
        _ = v18
        _ = v19
    } else {
        fmt.Println(v2, v2)
        v2 = true
        for fi20 := 0; fi20 < 8; fi20++ {
            v1 = append(v1, "hello")
            _ = fi20
        }
    }
    v4 = append(v4, 276)
    v3 = false
    var v21 uint64 = 2
    if (!v3) {
        v5 -= v5
        switch -1 {
        case 0:
            v2 = v2
        case -2147483648:
            v22 := GenS1{F0: 4294967295}
            _ = v22
        case 128:
            var v23 int32 = 42
            _ = v23
        default:
            var v24 int32 = 2147483647
            _ = v24
        }
        switch "\n\t" {
        case "\n\t":
            if (4294967295 != -9223372036854775808) {
                v21 = v21
                var v25 int32 = -1
                _ = v25
            }
        case "foo bar":
            for fi26 := 0; fi26 < 15; fi26++ {
                v2 = (!((!(!(v2 && (!((-32768 > -32768) || ((v3 && (v2 && v2)) || true)))))) && v3))
                _ = fi26
            }
        case "hello":
            switch -368 {
            case 9223372036854775807:
                v3 = v3
            case 4294967295:
                var v27 int16 = 73
                _ = v27
            }
        }
    }
    if (!v3) {
        for fi28 := 0; fi28 < 19; fi28++ {
            _ = len(v1)
            if false {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v1 = append(v1, ("hello" + "hello"))
                    var v30 uint8 = 127
                    var v31 uint64 = 18446744073709551615
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = len(v1)
                fmt.Println(fi28)
            } else {
                v32 := GenS1{F0: fi28}
                _ = v32
            }
            switch "hello" {
            case "":
                var v33 uint16 = v5
                _ = v33
            default:
                if v2 {
                    _ = len(v1)
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v34 uint64 = v21
                    v2 = false
                    _ = v34
                }
            }
            _ = fi28
        }
        if (v3 && v3) {
            var v35 int8 = -84
            for fi36 := 0; fi36 < 19; fi36++ {
                if len(v4) > 0 {
                    v4[0] = -1
                }
                _ = fi36
            }
            v37 := GenS2{F0: 1, F1: 32767, F2: 371}
            _ = v35
            _ = v37
        } else {
            for fi38 := 0; fi38 < 4; fi38++ {
                var v39 uint64 = v21
                var v40 uint64 = (v21 * 2)
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    }
    v21 = (v21 ^ 42)
    var v41 int64 = 42
    if (0 > 128) {
        var v42 int32 = -2147483648
        _ = v42
    }
    v43 := GenS3{F0: 127, F1: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
}
