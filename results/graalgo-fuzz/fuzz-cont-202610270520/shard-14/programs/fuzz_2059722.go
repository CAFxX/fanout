package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS4 struct { F0 uint64; F1 int64; F2 uint8 }
func genFn0() int8 {
    switch ("foo bar" + "a") {
    case "\n\t":
        var v7 bool = true
        _ = v7
    case "":
        var v8 uint16 = 65535
        _ = v8
    case "a":
        v9 := []int{2147483647}
        _ = v9
    default:
        v10 := []int{42, -447}
        _ = v10
    }
    var v11 uint = 2
    _ = v11
    return -75
}
func genFn1() uint8 {
    for fi12 := 0; fi12 < 14; fi12++ {
        if (!(!((1.5 - 1e10) > ((-2.25 + float32(1e10)) - 1.5)))) {
            for fi13 := 0; fi13 < 6; fi13++ {
                if true {
                    var v14 uint16 = 65535
                    v14 = 2
                    _ = v14
                }
                switch (-(fi13)) {
                case -730:
                    var v15 uint = 18446744073709551615
                    _ = v15
                case 9223372036854775807:
                    v16 := GenS1{F0: 65535}
                    _ = v16
                case 32768:
                    var v17 uint32 = 1
                    _ = v17
                default:
                    fmt.Println(fi12, fi12)
                }
                _ = fi13
            }
            if (false && true) {
                var v18 uint16 = 255
                fmt.Println(fi12)
                _ = v18
            }
            var v19 uint64 = 18446744073709551615
            _ = v19
        } else {
            v20 := GenS1{F0: 522}
            v21 := GenS2{F0: 0, F1: 2, F2: 1}
            _ = v20
            _ = v21
        }
        fmt.Println(fi12)
        _ = fi12
    }
    if (-128 >= genFn0()) {
        var v22 uint64 = 18446744073709551615
        var v23 string = "foo bar"
        v24 := []int{}
        _ = v22
        _ = v23
        _ = v24
    }
    v25 := GenS3{F0: 41, F1: 2, F2: 1}
    var v26 uint64 = 645
    _ = v25
    _ = v26
    return 244
}

func main() {
    v1 := []int{4294967295, -1, 0}
    v2 := make(map[string]uint32)
    var v3 float64 = ((0.1 - -2.25) / 3.14159)
    var v4 int32 = 0
    v5 := []string{"\n\t", ("hello" + ("hello" + "a")), "a", "foo bar"}
    var v6 uint16 = 20
    switch -9223372036854775808 {
    case -9223372036854775808:
        v6 = (v6 >> ((0) & 63))
    }
    switch 32767 {
    case -957:
        var v27 int32 = 864
        _ = v27
    case 128:
        var v28 uint32 = 0
        _ = v28
    case 406:
        var v29 float32 = (-2.25 + (0.1 - float32(v3)))
        _ = v29
    }
    if false {
        var v30 int8 = 63
        _ = v30
    }
    v31 := GenS3{F0: 209, F1: 42, F2: v4}
    if (2147483647 > 505) {
        var v32 int32 = 42
        var v33 uint32 = 30
        _ = v32
        _ = v33
    }
    v34 := GenS1{F0: 521}
    var v35 int = 1
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 float32 = float32(float64(-2.25))
        _ = fi36
        _ = v37
    }
    if (4294967295 != 42) {
        if (!((false && true) || false)) {
            for fi38 := 0; fi38 < 20; fi38++ {
                v39 := GenS3{F0: 189, F1: 128, F2: v4}
                fmt.Println(v6)
                _ = fi38
                _ = v39
            }
            v5 = append(v5, "hello")
            var v40 uint32 = 4294967295
            _ = v40
        }
        var v41 bool = (-2147483648 >= v4)
        _ = v41
    } else {
        if false {
            var v42 int = (v35 - (v35 << ((0) & 63)))
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v43 uint8 = 128
            _ = v42
            _ = v43
        }
        _ = len(v2)
        fmt.Println(v4, v3)
    }
    if false {
        v44 := GenS4{F0: 2, F1: 606, F2: genFn1()}
        v45 := GenS4{F0: 418, F1: 1, F2: 2}
        v6 = v6
        _ = v44
        _ = v45
    } else {
        v4 = v4
        fmt.Println(v4, v3)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
