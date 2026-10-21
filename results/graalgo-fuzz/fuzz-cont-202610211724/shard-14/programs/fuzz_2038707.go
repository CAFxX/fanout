package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 uint64) int32 {
    p0 += (p0 - (p0 * (p0 & (p0 * (p0 - (p0 << ((4294967295) & 63)))))))
    var v6 string = "hello"
    var v7 string = v6
    if true {
        if ((255 <= 32767) || true) {
            v8 := GenS1{F0: 128, F1: 1, F2: 2}
            _ = v8
        }
        v9 := []int{128, 32768}
        _ = v9
    } else {
        var v10 bool = false
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 uint16 = 65535
    v2 := []string{"a", ("a" + (("hello" + "\n\t") + "")), "hello", "\n\t"}
    var v3 uint16 = (v1 - (v1 ^ (v1 + v1)))
    var v4 int32 = 0
    v5 := make(map[string]string)
    fmt.Println(v4)
    for fi11 := 0; fi11 < 4; fi11++ {
        if false {
            if (true || true) {
                var v12 int8 = 100
                _ = v12
            }
            if (-2.25 > ((3.14159 * (1.5 / (-2.25 + 1e10))) / (3.14159 / (0.1 / 1.0)))) {
                var v13 int = -9223372036854775808
                var v14 float64 = -2.25
                fmt.Println(v4, v4, v14)
                _ = v13
                _ = v14
            }
            v4 += (v4 ^ v4)
        }
        v15 := GenS2{F0: 100, F1: 1, F2: 0}
        if (((1.5 * 3.14159) > float64(-2.25)) || false) {
            v5[(("a" + "foo bar") + ("hello" + ""))] = "hello"
            var v16 string = ""
            _ = v16
        } else {
            v17 := GenS3{F0: 2}
            _ = v17
        }
        _ = fi11
        _ = v15
    }
    var v18 int64 = 42
    var v19 float64 = 1e10
    switch int(v18) {
    case 604:
        v18 = v18
    case -9223372036854775808:
        var v20 bool = (true || false)
        _ = v20
    }
    fmt.Println(v1, v4, v18)
    if (!true) {
        v21 := GenS4{F0: 0, F1: (^(v18))}
        v22 := GenS3{F0: 913}
        var v23 uint16 = 0
        _ = v21
        _ = v22
        _ = v23
    }
    var v24 uint64 = 1
    v3 -= 32767
    var v25 uint8 = 255
    v26 := GenS5{F0: ("hello" + "hello"), F1: v1, F2: 65535}
    var v27 uint16 = 65535
    for fi28 := 0; fi28 < 11; fi28++ {
        v18 -= (v18 * (-(v18)))
        _ = fi28
    }
    v24 -= 2
    v5["hello"] = "a"
    for fi29 := 0; fi29 < 3; fi29++ {
        if true {
            if (255 < v25) {
                if true {
                    if (!false) {
                        v30 := GenS4{F0: 42, F1: (v18 << ((65535) & 63))}
                        var v31 bool = ((v4 <= (^(v4))) && (true || (335 != (^(v1)))))
                        fmt.Println(v19, v4, v4)
                        _ = v30
                        _ = v31
                    } else {
                        var v32 uint32 = 65535
                        _ = v32
                    }
                } else {
                    v18 = (v18 % (((^(v18))) | 1))
                }
                var v33 uint32 = 1
                _ = v33
            }
            v34 := GenS1{F0: (v25 ^ v25), F1: uint32(v24), F2: uint(v24)}
            _ = v34
        } else {
            v35 := v5["a"]
            fmt.Println(fi29, v27, v19)
            var v36 int = (fi29 / ((32767) | 1))
            _ = v35
            _ = v36
        }
        v25 += (v25 - uint8(v3))
        for fi37 := 0; fi37 < 20; fi37++ {
            switch "" {
            case "\n\t":
                fmt.Println(fi29)
            case "":
                if false {
                    var v38 int64 = v18
                    _ = v38
                }
            }
            if (!((v3 * (v1 & (v3 / ((v1) | 1)))) < (^(v3)))) {
                v4 = int32(v25)
                var v39 uint32 = uint32(v24)
                var v40 uint64 = (v24 % (((^(v24))) | 1))
                _ = v39
                _ = v40
            } else {
                fmt.Println(v18, fi29, v1)
                var v41 uint64 = v24
                fmt.Println(v19, v18, fi37)
                _ = v41
            }
            _ = fi37
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
