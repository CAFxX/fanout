package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 string, p1 int64) uint8 {
    switch "foo bar" {
    case "\n\t":
        for fi7 := 0; fi7 < 10; fi7++ {
            if true {
                p0 = ("hello" + "\n\t")
                v8 := []int{(-(fi7))}
                if (1 >= 2) {
                    p1 -= (p1 >> ((65535) & 63))
                } else {
                    var v9 int32 = 282
                    _ = v9
                }
                _ = v8
            } else {
                var v10 int = (fi7 % (((-(fi7))) | 1))
                _ = v10
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                if (398 != 0) {
                    var v12 int = -650
                    v13 := []int{(fi11 + 128), -9223372036854775808, v12}
                    _ = len(v13)
                    _ = v12
                    _ = v13
                } else {
                    fmt.Println(fi7)
                    v14 := GenS1{F0: "foo bar", F1: 32767, F2: (fi7 << ((0) & 63))}
                    _ = v14
                }
                p0 = "\n\t"
                v15 := GenS2{F0: 1, F1: 127}
                _ = fi11
                _ = v15
            }
            var v16 int32 = 420
            _ = fi7
            _ = v16
        }
    case "":
        v17 := []int{}
        _ = v17
    case "a":
        switch 2147483647 {
        case 128:
            var v18 int8 = -9
            _ = v18
        case -300:
            v19 := GenS2{F0: 76, F1: 42}
            _ = v19
        }
    default:
        var v20 int16 = 31
        _ = v20
    }
    switch ((((p0 + "a") + ((p0 + p0) + p0)) + p0) + p0) {
    case "hello":
        var v21 uint8 = 128
        _ = v21
    }
    v22 := GenS3{F0: -2147483648}
    _ = p0
    _ = p1
    _ = v22
    return 255
}

func main() {
    var v1 uint64 = 259
    var v2 uint16 = 65535
    var v3 int16 = -955
    v4 := make(map[string]uint32)
    var v5 int32 = -269
    v6 := make(map[int]uint32)
    v5 = (v5 % (((v5 << ((1) & 63))) | 1))
    var v23 uint64 = (^(v1))
    if true {
        switch 4294967295 {
        case 32767:
            v1 = (v23 / ((580) | 1))
        case 9223372036854775807:
            var v24 string = "a"
            _ = v24
        default:
            for fi25 := 0; fi25 < 7; fi25++ {
                v1 = 0
                v26 := []int{}
                _ = fi25
                _ = v26
            }
        }
        var v27 int8 = 100
        _ = v27
    }
    v23 += v23
    fmt.Println(v2, v23, v5)
    v1 = (v23 / ((v1) | 1))
    switch 32767 {
    case 32767:
        delete(v6, 1)
    case -1:
        for fi28 := 0; fi28 < 11; fi28++ {
            if true {
                switch fi28 {
                case -233:
                    if ("\n\t" != "a") {
                        for fi29 := 0; fi29 < 9; fi29++ {
                            var v30 uint8 = 124
                            var v31 bool = (true || (0 <= v3))
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    } else {
                        v2 = (v2 % ((v2) | 1))
                    }
                case 42:
                    var v32 int32 = v5
                    _ = v32
                default:
                    var v33 int16 = (v3 >> ((1) & 63))
                    _ = v33
                }
                v34 := GenS3{F0: v5}
                _ = v34
            } else {
                for fi35 := 0; fi35 < 17; fi35++ {
                    var v36 int64 = 1
                    _ = fi35
                    _ = v36
                }
            }
            v2 = v2
            _ = fi28
        }
    }
    if ((v23 * 487) != (^(v1))) {
        fmt.Println(v3, v2, v3)
    } else {
        var v37 uint8 = 128
        _ = v37
    }
    var v38 int = 1
    v5 = v5
    v23 = 18446744073709551615
    v39 := []int{v38, 740}
    for fi40 := 0; fi40 < 17; fi40++ {
        var v41 bool = (0 <= 352)
        _ = fi40
        _ = v41
    }
    var v42 uint16 = (v2 | v2)
    v43 := GenS1{F0: "foo bar", F1: v3, F2: v38}
    fmt.Println(v5)
    _ = len(v39)
    var v44 uint16 = v42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
