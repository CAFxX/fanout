package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 uint; F1 uint; F2 string }
type GenS5 struct { F0 string }
func genFn0(p0 int16, p1 uint, p2 uint8) string {
    switch ("hello" + "a") {
    case "":
        v5 := []int{-360, 32768}
        _ = v5
    case "a":
        if true {
            v6 := GenS1{F0: -1, F1: 615}
            _ = v6
        }
    }
    p1 = (p1 & 393)
    p1 = (p1 << ((p1) & 63))
    for fi7 := 0; fi7 < 13; fi7++ {
        if true {
            fmt.Println(p2, p1)
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]bool)
    var v3 uint = 65535
    var v4 bool = false
    fmt.Println(v3, v4)
    if v4 {
        _ = len(v2)
    }
    v3 = 4294967295
    if false {
        v8 := GenS2{F0: 2147483647, F1: 663, F2: -9223372036854775808}
        _ = v8
    } else {
        if (2 > 866) {
            var v9 uint = v3
            _ = v9
        }
    }
    if (!v4) {
        switch -1 {
        case 32767:
            var v10 float32 = (3.14159 * (float32(3.14159) - (0.0 - (0.1 / 0.1))))
            _ = v10
        default:
            fmt.Println(v3, v4)
        }
    } else {
        if v4 {
            var v11 int64 = 1
            var v12 int = -732
            _ = v11
            _ = v12
        } else {
            v3 = (v3 ^ v3)
            for fi13 := 0; fi13 < 10; fi13++ {
                if (v4 && v4) {
                    fmt.Println(v4)
                    var v14 float64 = 0.1
                    for fi15 := 0; fi15 < 1; fi15++ {
                        v16 := GenS3{F0: 135, F1: 42}
                        if ((-953 >= 0) || (!v4)) {
                            v17 := GenS1{F0: 1, F1: 1}
                            v18 := GenS1{F0: -1, F1: 1}
                            _ = v17
                            _ = v18
                        } else {
                            v14 += v14
                        }
                        _ = fi15
                        _ = v16
                    }
                    _ = v14
                } else {
                    v19 := []int{(fi13 | fi13), (fi13 & fi13)}
                    _ = v19
                }
                if (100 != -52) {
                    v4 = (-1 > 0)
                    v4 = (1e10 == 1e10)
                    v20 := []int{}
                    _ = v20
                }
                _ = fi13
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                var v22 float64 = ((-2.25 * (0.1 + 3.14159)) - float64((0.0 * 0.0)))
                v4 = v4
                _ = fi21
                _ = v22
            }
        }
        v4 = false
    }
    for fi23 := 0; fi23 < 12; fi23++ {
        if v4 {
            var v24 int64 = 0
            v4 = (18446744073709551615 == (v3 + v3))
            _ = v24
        } else {
            v3 -= 11
        }
        var v25 uint8 = 35
        v26 := []int{fi23, (fi23 ^ fi23), fi23}
        _ = fi23
        _ = v25
        _ = v26
    }
    v27 := v2[("foo bar" + "\n\t")]
    var v28 uint8 = 42
    switch -993 {
    case 9223372036854775807:
        v3 += (v3 + 0)
    case 1:
        v3 -= 686
    default:
        v27 = (false && (v4 && v4))
    }
    if v4 {
        v29 := []int{-2147483648}
        v27 = (!(!true))
        _ = v29
    } else {
        v4 = (!true)
        for fi30 := 0; fi30 < 9; fi30++ {
            v31 := []int{32767, (fi30 & fi30)}
            _ = fi30
            _ = v31
        }
        if v4 {
            switch "\n\t" {
            case "a":
                v4 = (100 != 127)
            }
            var v32 int32 = 952
            v33 := GenS3{F0: (^(v28)), F1: -9223372036854775808}
            _ = v32
            _ = v33
        } else {
            var v34 uint8 = (v28 % ((v28) | 1))
            v27 = (127 < 127)
            _ = len(v1)
            _ = v34
        }
    }
    var v35 int = 4294967295
    var v36 uint64 = 88
    if (!(2 >= 893)) {
        v37 := GenS4{F0: v3, F1: (v3 >> ((v3) & 63)), F2: ("\n\t" + ("a" + "a"))}
        _ = v37
    } else {
        var v38 uint16 = 637
        fmt.Println(v35, v36, v36)
        _ = v38
    }
    if v4 {
        v39 := GenS5{F0: ("hello" + (("hello" + "\n\t") + "foo bar"))}
        switch "" {
        case "a":
            v36 += v36
        case "\n\t":
            var v40 float64 = 0.1
            _ = v40
        default:
            for fi41 := 0; fi41 < 1; fi41++ {
                var v42 float64 = -2.25
                _ = fi41
                _ = v42
            }
        }
        var v43 string = "a"
        _ = v39
        _ = v43
    } else {
        fmt.Println(v36)
        var v44 uint64 = 4294967295
        v45 := GenS4{F0: (v3 % (((v3 << ((v3) & 63))) | 1)), F1: (v3 ^ (v3 - v3)), F2: "hello"}
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v36
    _ = v4
}
