package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS3 struct { F0 int32; F1 int64 }
func genFn0(p0 int64, p1 int16) uint {
    var v6 bool = true
    v6 = ((p0 / ((p0) | 1)) == -1)
    p0 -= 0
    var v7 bool = true
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]int)
    var v3 uint64 = 243
    var v4 uint32 = 42
    var v5 int8 = 1
    if true {
        v5 = int8(v4)
        switch 0 {
        case 4294967295:
            fmt.Println(v5, v4, v5)
        case 1:
            if true {
                v5 = (v5 & (^(v5)))
                fmt.Println(v3)
            }
        default:
            delete(v2, "hello")
        }
    }
    var v8 uint = 1
    if true {
        v9 := GenS1{F0: (v4 >> ((v8) & 63)), F1: 1}
        _ = v9
    } else {
        var v10 int64 = -399
        _ = v10
    }
    var v11 float32 = ((0.1 / 1e10) / 1e10)
    switch ("foo bar" + ("foo bar" + "")) {
    case "\n\t":
        if true {
            for fi12 := 0; fi12 < 18; fi12++ {
                switch ("" + (("a" + "hello") + ("\n\t" + "foo bar"))) {
                case "hello":
                    for fi13 := 0; fi13 < 11; fi13++ {
                        v14 := GenS1{F0: 42, F1: v3}
                        v3 = (v3 << ((genFn0(1, -1)) & 63))
                        _ = fi13
                        _ = v14
                    }
                case "a":
                    var v15 bool = ((^(v3)) < (v3 ^ (v3 % ((387) | 1))))
                    _ = v15
                case "\n\t":
                    fmt.Println(v8, v4)
                }
                _ = fi12
            }
            v16 := GenS2{F0: 42, F1: 0, F2: uint64(v4)}
            _ = v16
        }
    case "":
        v5 = v5
    case "a":
        fmt.Println(v11, v8)
    }
    v11 -= (0.0 - v11)
    if false {
        fmt.Println(v3, v8, v11)
        var v17 uint32 = v4
        var v18 string = "\n\t"
        _ = v17
        _ = v18
    }
    v19 := GenS1{F0: 4294967295, F1: v3}
    for fi20 := 0; fi20 < 3; fi20++ {
        delete(v1, ("" + "a"))
        var v21 float32 = v11
        v5 = (-(v5))
        _ = fi20
        _ = v21
    }
    var v22 bool = (255 <= 128)
    for fi23 := 0; fi23 < 14; fi23++ {
        switch (("a" + "\n\t") + "hello") {
        case "a":
            if (344 == 32767) {
                for fi24 := 0; fi24 < 19; fi24++ {
                    v11 += 0.1
                    if (fi23 != (fi23 << (((v8 * v8)) & 63))) {
                        var v25 float32 = ((1.5 + 0.0) * (v11 * (v11 * 0.0)))
                        _ = v25
                    } else {
                        v26 := GenS1{F0: 1, F1: (v3 + 944)}
                        _ = v26
                    }
                    _ = fi24
                }
                if false {
                    if (!(((uint(v3) < v8) || v22) || v22)) {
                        delete(v1, (((("hello" + "") + "") + (("a" + ("\n\t" + ("\n\t" + (("hello" + "\n\t") + "")))) + "hello")) + "a"))
                        for fi27 := 0; fi27 < 8; fi27++ {
                            if false {
                                var v28 int8 = (-(v5))
                                v29 := v2["foo bar"]
                                _ = v28
                                _ = v29
                            } else {
                                var v30 uint8 = 42
                                _ = v30
                            }
                            v5 -= (v5 * -108)
                            var v31 int64 = -1
                            _ = fi27
                            _ = v31
                        }
                        v11 = 1e10
                    }
                }
            }
        case "hello":
            var v32 uint32 = (v4 * (^(v4)))
            _ = v32
        default:
            if ((v3 * v3) > (v3 * v3)) {
                if false {
                    var v33 bool = ("a" > ("a" + ""))
                    var v34 string = "foo bar"
                    v34 = v34
                    _ = v33
                    _ = v34
                }
                var v35 int16 = 32767
                _ = v35
            } else {
                v36 := v2["a"]
                v37 := GenS3{F0: 42, F1: 287}
                _ = v36
                _ = v37
            }
        }
        _ = fi23
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        v5 = v5
        var v39 string = "hello"
        _ = fi38
        _ = v39
    }
    var v40 bool = (true || v22)
    if (v11 != (1e10 / -2.25)) {
        var v41 uint8 = 42
        var v42 string = "foo bar"
        _ = v41
        _ = v42
    } else {
        fmt.Println(v3)
        var v43 uint64 = v3
        v44 := []int{532, 523, -322, 504}
        _ = v43
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v8
}
