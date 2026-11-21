package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int64; F1 uint32; F2 bool }
type GenS4 struct { F0 int; F1 int32; F2 uint64 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint64, p1 bool) int16 {
    if (p1 || false) {
        v5 := GenS1{F0: 0}
        fmt.Println(p0, p0)
        p1 = (!p1)
        _ = v5
    } else {
        var v6 int64 = -758
        v6 = (v6 - (^(v6)))
        _ = v6
    }
    v7 := []int{614, 822, 0, 32768}
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 uint) uint {
    v8 := []int{2147483647, -1}
    var v9 string = ("\n\t" + "foo bar")
    p1 += 4294967295
    if true {
        v9 = v9
        var v10 uint = p1
        switch ((((("foo bar" + "") + v9) + "foo bar") + v9) + v9) {
        case "":
            v8 = append(v8, -2147483648)
        default:
            for fi11 := 0; fi11 < 9; fi11++ {
                var v12 uint64 = 4294967295
                if (!(42 < 42)) {
                    if ((((127 > 1) && (!(!true))) && true) && (0 >= 773)) {
                        switch v9 {
                        case "a":
                            if len(v8) > 0 {
                                v8[0] = (^(fi11))
                            }
                        case "":
                            fmt.Println(v10)
                        case "foo bar":
                            v10 -= 255
                        default:
                            switch v9 {
                            case "hello":
                                v13 := GenS2{F0: ((v9 + (v9 + "foo bar")) + v9), F1: -1}
                                _ = v13
                            case "":
                                _ = len(v8)
                            default:
                                switch 128 {
                                case -1:
                                    fmt.Println(v9, p0)
                                case -313:
                                    var v14 uint32 = 1
                                    _ = v14
                                case 4294967295:
                                    v15 := GenS3{F0: 42, F1: 312, F2: true}
                                    _ = v15
                                }
                            }
                        }
                        if false {
                            var v16 int64 = 0
                            for fi17 := 0; fi17 < 12; fi17++ {
                                p1 -= (v10 >> ((2) & 63))
                                _ = fi17
                            }
                            _ = v16
                        }
                    }
                } else {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        fmt.Println(p0, v10, fi18)
                        v9 = (v9 + "")
                        v19 := GenS1{F0: -2147483648}
                        _ = fi18
                        _ = v19
                    }
                    _ = len(v8)
                }
                _ = fi11
                _ = v12
            }
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (p1 >> ((230) & 63))
}

func main() {
    v1 := []string{("\n\t" + ""), ((("foo bar" + "\n\t") + ("a" + "")) + "\n\t"), ("a" + ((((("" + "foo bar") + ("" + "foo bar")) + ("" + "\n\t")) + "") + (("\n\t" + (("\n\t" + (("foo bar" + "hello") + "a")) + "foo bar")) + ("" + "\n\t")))), "a"}
    var v2 float64 = float64(0.1)
    var v3 int = -2147483648
    var v4 float64 = v2
    v3 = v3
    var v20 int8 = 6
    fmt.Println(v20, v4)
    v21 := GenS4{F0: 820, F1: 2147483647, F2: 929}
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "hello")
    }
    var v22 int = v3
    var v23 uint = 18446744073709551615
    v24 := GenS5{F0: (42 > 1)}
    v20 = (^(v20))
    for fi25 := 0; fi25 < 19; fi25++ {
        if (false && (int(v23) < fi25)) {
            var v26 uint8 = 97
            _ = v26
        }
        var v27 float64 = v2
        fmt.Println(v27, v4)
        _ = fi25
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
