package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 string }
func genFn0(p0 string, p1 int16, p2 int64) uint32 {
    for fi7 := 0; fi7 < 4; fi7++ {
        if (!(846 > 65535)) {
            v8 := GenS1{F0: 2147483647}
            _ = v8
        } else {
            var v9 bool = (-529 == -1)
            v10 := []int{(fi7 / ((fi7) | 1)), fi7}
            fmt.Println(p2)
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    p2 = 0
    _ = p0
    _ = p1
    _ = p2
    return 995
}
func genFn1() string {
    switch -1 {
    case -9223372036854775808:
        var v11 int8 = -1
        _ = v11
    case 290:
        if (!(!(128 != 9223372036854775807))) {
            var v12 uint = 0
            v12 += v12
            switch ((("" + "foo bar") + "hello") + "foo bar") {
            case "":
                var v13 int16 = 1000
                _ = v13
            default:
                var v14 int8 = -43
                _ = v14
            }
            _ = v12
        }
    case -555:
        var v15 int32 = -153
        _ = v15
    }
    var v16 int = 9223372036854775807
    v16 -= v16
    fmt.Println(v16)
    _ = v16
    return "hello"
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := make(map[int]string)
    var v3 uint = 438
    v4 := make(map[uint32]uint32)
    var v5 bool = ("hello" < ("\n\t" + ("\n\t" + ("a" + "\n\t"))))
    var v6 int16 = -32768
    v6 -= (v6 | 269)
    delete(v2, 859)
    v17 := GenS1{F0: 863}
    v18 := GenS1{F0: 4294967295}
    var v19 uint32 = 42
    v20 := []int{32767}
    for fi21 := 0; fi21 < 12; fi21++ {
        if ((!true) || ((fi21 & (fi21 % ((962) | 1))) == (-(fi21)))) {
            v3 -= (v3 + 0)
        }
        var v22 uint64 = 1
        if (-70 != 100) {
            for fi23 := 0; fi23 < 19; fi23++ {
                v5 = ("\n\t" > "hello")
                var v24 uint32 = v19
                v25 := GenS2{F0: -1, F1: (v3 << (((v3 << ((v3) & 63))) & 63)), F2: 255}
                _ = fi23
                _ = v24
                _ = v25
            }
            v6 -= -32768
        } else {
            v22 = v22
            var v26 uint64 = 42
            _ = v26
        }
        _ = fi21
        _ = v22
    }
    var v27 int32 = 128
    for fi28 := 0; fi28 < 3; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            if (-595 <= v6) {
                var v30 int = fi28
                _ = v30
            } else {
                v31 := GenS3{F0: "hello"}
                v2[127] = ((("a" + "") + "\n\t") + "\n\t")
                _ = v31
            }
            v6 = (v6 & v6)
            v27 -= v27
            if len(v20) > 0 {
                v20[0] = fi29
            }
            var v32 float64 = -2.25
            _ = fi29
            _ = v32
        }
        v33 := GenS2{F0: 786, F1: (v3 ^ (v3 * (v3 ^ 43))), F2: 255}
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi28
        _ = v33
    }
    if (!false) {
        v34 := GenS2{F0: (v27 / ((v27) | 1)), F1: v3, F2: 153}
        if (18446744073709551615 == 110) {
            var v35 int64 = (v1 << ((v3) & 63))
            if true {
                if (949 > 40) {
                    var v36 uint8 = 2
                    fmt.Println(v27, v27, v1)
                    v4[v19] = (v19 % ((v19) | 1))
                    _ = v36
                } else {
                    var v37 int = -220
                    _ = v37
                }
                var v38 int64 = v1
                _ = v38
            }
            if v5 {
                v39 := GenS2{F0: v27, F1: v3, F2: 2}
                v40 := GenS4{F0: 32767, F1: "\n\t"}
                _ = v39
                _ = v40
            } else {
                v5 = (!true)
                v19 += (v19 - (v19 | v19))
            }
            _ = v35
        } else {
            switch -1 {
            case -426:
                v20 = append(v20, 478)
            }
            v3 = (v3 * v3)
            fmt.Println(v1, v1, v1)
        }
        _ = v34
    } else {
        v4[(^(v19))] = v19
        var v41 uint8 = 236
        v42 := GenS1{F0: (v19 + genFn0("", 1000, v1))}
        _ = v41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
