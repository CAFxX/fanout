package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
func genFn0() uint32 {
    var v5 int16 = -412
    v6 := GenS1{F0: -1}
    fmt.Println(v5)
    var v7 bool = ((!(65535 >= 1)) || true)
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 uint16) uint32 {
    var v8 uint = 743
    fmt.Println(p0, p0)
    var v9 uint16 = 42
    fmt.Println(v9)
    _ = p0
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{244, 2147483648}
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    v4 := []int{0, 566}
    v10 := GenS1{F0: -1}
    var v11 int32 = 1
    var v12 uint8 = 128
    fmt.Println(v12)
    fmt.Println(v11, v12)
    if (!(0 != 0)) {
        v11 += v11
        var v13 string = (("\n\t" + ("\n\t" + (((((("foo bar" + "hello") + "") + "") + "foo bar") + (("hello" + "") + "")) + "\n\t"))) + ("" + "foo bar"))
        if (((950 >= -2147483648) || (((0.1 - 0.0) + float32(0.0)) > 0.0)) && (((-2.25 == 1e10) && (("hello" < "\n\t") || (true && false))) || (v12 < (v12 % (((^(v12))) | 1))))) {
            v14 := GenS2{F0: 0}
            v11 = v11
            _ = v14
        } else {
            if (true || true) {
                var v15 bool = true
                if (841 != 962) {
                    v15 = (-2.25 <= 1e10)
                    fmt.Println(v13, v12)
                } else {
                    var v16 float64 = -2.25
                    _ = v16
                }
                var v17 bool = (false && v15)
                _ = v15
                _ = v17
            }
            v12 = (v12 & v12)
            fmt.Println(v12, v11, v13)
        }
        _ = v13
    } else {
        if true {
            switch "a" {
            case "\n\t":
                if true {
                    v18 := GenS3{F0: v12}
                    _ = v18
                } else {
                    switch "\n\t" {
                    case "hello":
                        v11 += v11
                    case "\n\t":
                        v12 -= 1
                    default:
                        var v19 bool = false
                        _ = v19
                    }
                    if len(v4) > 0 {
                        v4[0] = 128
                    }
                }
            case "foo bar":
                if false {
                    var v20 float32 = 1e10
                    _ = len(v4)
                    v20 -= v20
                    _ = v20
                } else {
                    v11 = v11
                }
            default:
                var v21 int16 = 0
                _ = v21
            }
        }
        switch -9223372036854775808 {
        case 2147483647:
            var v22 uint64 = 4294967295
            _ = v22
        }
        fmt.Println(v11, v12)
    }
    var v23 uint = 4294967295
    if (42 < 0) {
        if len(v1) > 0 {
            v1[0] = 2147483648
        }
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 8; fi25++ {
            v23 = v23
            var v26 string = "foo bar"
            _ = fi25
            _ = v26
        }
        _ = fi24
    }
    v27 := GenS2{F0: -128}
    var v28 int32 = (v11 / ((0) | 1))
    switch 4294967295 {
    case 32768:
        var v29 string = "\n\t"
        _ = v29
    case 684:
        v12 += v12
    case -436:
        switch -837 {
        case 663:
            for fi30 := 0; fi30 < 14; fi30++ {
                switch ("" + ("foo bar" + "a")) {
                case "":
                    v31 := GenS1{F0: -32768}
                    _ = v31
                case "a":
                    fmt.Println(v23, v23, fi30)
                default:
                    var v32 int8 = -1
                    _ = v32
                }
                _ = fi30
            }
        case 4294967295:
            v11 = -434
        case 128:
            var v33 uint64 = 269
            _ = v33
        }
    }
    fmt.Println(v23, v12)
    fmt.Println(v11, v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
