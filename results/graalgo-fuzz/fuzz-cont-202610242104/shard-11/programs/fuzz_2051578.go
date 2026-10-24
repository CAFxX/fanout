package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int16, p1 uint, p2 string) uint8 {
    v7 := GenS1{F0: 255, F1: 0, F2: (false && (p2 <= p2))}
    p2 = ("hello" + "foo bar")
    var v8 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int16, p1 int64, p2 int32) bool {
    for fi9 := 0; fi9 < 19; fi9++ {
        if true {
            if true {
                p2 = (^(p2))
            }
            p1 = (p1 / (((p1 - p1)) | 1))
            for fi10 := 0; fi10 < 6; fi10++ {
                switch "foo bar" {
                case "\n\t":
                    if true {
                        p1 += (p1 & 42)
                    } else {
                        var v11 int64 = 805
                        _ = v11
                    }
                case "":
                    var v12 float64 = 0.0
                    _ = v12
                }
                var v13 uint64 = 42
                v14 := GenS1{F0: 192, F1: v13, F2: false}
                _ = fi10
                _ = v13
                _ = v14
            }
        } else {
            v15 := []int{}
            _ = v15
        }
        v16 := []int{fi9, fi9, fi9}
        if (-1 == -400) {
            fmt.Println(p2, p1)
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 uint16 = 1
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                p1 = int64(fi17)
                _ = fi17
                _ = v18
            }
            var v19 uint = 2
            _ = v19
        } else {
            if false {
                v20 := GenS2{F0: 42, F1: 1}
                _ = v20
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 bool = (p1 < 9223372036854775807)
                if v22 {
                    _ = len(v16)
                    if (fi9 != (fi9 / ((9223372036854775807) | 1))) {
                        p2 += p2
                    } else {
                        p0 += p0
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        p1 = p1
                        p2 += (p2 | 2147483647)
                    }
                    var v23 int32 = p2
                    _ = v23
                } else {
                    var v24 int32 = p2
                    var v25 int = 9223372036854775807
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        }
        _ = fi9
        _ = v16
    }
    p2 -= (^(p2))
    v26 := GenS2{F0: 751, F1: 255}
    p2 -= p2
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return ((((("foo bar" + ("a" + ("hello" + "\n\t"))) + ("\n\t" + "a")) + "a") + "hello") >= ("\n\t" + ("hello" + (("\n\t" + ("hello" + ("a" + ("" + "a")))) + ("" + "hello")))))
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]int)
    var v3 int8 = -1
    var v4 int64 = -312
    v5 := make(map[string]int)
    var v6 uint8 = 232
    var v27 uint = 65535
    for fi28 := 0; fi28 < 17; fi28++ {
        if false {
            v29 := []int{fi28, (^(fi28)), 127, (-(fi28))}
            fmt.Println(fi28)
            _ = v29
        } else {
            var v30 uint16 = 255
            _ = v30
        }
        var v31 int8 = v3
        v32 := []int{9223372036854775807, (fi28 % (((fi28 >> (((^(v27))) & 63))) | 1))}
        _ = fi28
        _ = v31
        _ = v32
    }
    v1 = genFn0(1, v27, "a")
    var v33 uint32 = 65535
    fmt.Println(v33, v27)
    var v34 int = 32767
    var v35 int32 = 1
    var v36 int64 = 42
    v27 -= 545
    var v37 int32 = (-(v35))
    v4 = -1
    for fi38 := 0; fi38 < 4; fi38++ {
        var v39 float32 = 0.1
        _ = fi38
        _ = v39
    }
    if false {
        if true {
            v4 = -9223372036854775808
            var v40 uint8 = (v1 + genFn0(1, uint(v33), ""))
            _ = v40
        }
    } else {
        v41 := []int{(v34 | (v34 & 396)), (v34 ^ 546), 718}
        _ = v41
    }
    if false {
        if true {
            switch 217 {
            case -2147483648:
                v4 -= (^(v36))
            case 2147483647:
                v42 := GenS1{F0: 13, F1: 482, F2: ((1.5 >= (0.0 / (0.1 + 0.1))) || false)}
                _ = v42
            case 42:
                v27 += (v27 << (((v27 << ((v27) & 63))) & 63))
            }
            v6 = 0
            var v43 int16 = 42
            _ = v43
        }
        fmt.Println(v6, v34, v34)
        switch "" {
        case "":
            v3 += v3
        }
    }
    var v44 float64 = float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
