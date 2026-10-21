package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int64; F1 string }
func genFn0() int {
    v6 := []int{2147483647, 983, -426}
    var v7 int = 9223372036854775807
    _ = v6
    _ = v7
    return (v7 | v7)
}
func genFn1(p0 int8, p1 uint16) string {
    p1 += p1
    if (0 <= 1) {
        v8 := GenS1{F0: (127 > 128), F1: 0}
        v9 := []int{32768, 1}
        _ = v8
        _ = v9
    }
    var v10 uint64 = 4294967295
    fmt.Println(p1, p1)
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 uint32 = 0
        if false {
            v10 = 893
            if true {
                for fi13 := 0; fi13 < 18; fi13++ {
                    var v14 uint64 = (v10 | uint64(p1))
                    switch (fi11 * genFn0()) {
                    case 788:
                        v15 := []int{fi13, (fi13 | fi11), fi13, (^(fi11))}
                        _ = v15
                    case 9223372036854775807:
                        if ((p1 % (((p1 / (((^(p1))) | 1))) | 1)) != (p1 ^ (p1 * (p1 >> ((2) & 63))))) {
                            v16 := []int{(fi13 - 1), fi11, fi13, fi11}
                            fmt.Println(v14)
                            _ = v16
                        }
                    default:
                        var v17 int64 = -17
                        _ = v17
                    }
                    p1 = p1
                    _ = fi13
                    _ = v14
                }
            }
            if true {
                var v18 uint64 = 1
                _ = v18
            } else {
                v12 = v12
                v19 := GenS1{F0: (!(false && true)), F1: v10}
                switch fi11 {
                case 0:
                    v12 = (v12 - 1)
                }
                _ = v19
            }
        }
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    return "foo bar"
}

func main() {
    var v1 uint = 866
    v2 := []int{2147483647, 9223372036854775807, -1}
    var v3 int32 = 1
    v4 := make(map[string]bool)
    v5 := make(map[string]int)
    var v20 string = ("hello" + "foo bar")
    var v21 uint32 = 1
    v20 = v20
    for fi22 := 0; fi22 < 17; fi22++ {
        if (((((v20 + "hello") + (v20 + ("" + v20))) + v20) + ((v20 + ("a" + v20)) + v20)) > v20) {
            if false {
                var v23 bool = ((v3 >> (((v1 % ((671) | 1))) & 63)) != (^(v3)))
                switch (v20 + (((v20 + v20) + v20) + "")) {
                case "\n\t":
                    if (!(v23 || ((-1 == v3) || ((!false) && v23)))) {
                        for fi24 := 0; fi24 < 8; fi24++ {
                            if ((!(!(128 != 42))) || true) {
                                if true {
                                    var v25 uint64 = 4294967295
                                    v2 = append(v2, (^(fi24)))
                                    _ = v25
                                }
                            }
                            var v26 string = v20
                            _ = fi24
                            _ = v26
                        }
                        fmt.Println(v21)
                    }
                case "foo bar":
                    var v27 int32 = (v3 << (((v1 & v1)) & 63))
                    _ = v27
                default:
                    var v28 int64 = 550
                    _ = v28
                }
                _ = v23
            }
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            if true {
                v30 := GenS1{F0: false, F1: 0}
                _ = v30
            } else {
                if (true && false) {
                    _ = len(v2)
                } else {
                    v21 -= (^(v21))
                }
            }
            if false {
                if (!("hello" < "foo bar")) {
                    var v31 uint16 = 0
                    var v32 uint64 = 42
                    _ = v31
                    _ = v32
                }
                if true {
                    if false {
                        var v33 int = (fi22 / (((fi22 << ((v1) & 63))) | 1))
                        _ = v33
                    }
                } else {
                    var v34 uint = v1
                    var v35 uint = v1
                    v34 = v1
                    _ = v34
                    _ = v35
                }
                _ = len(v5)
            } else {
                v36 := v4[v20]
                var v37 uint8 = 255
                _ = v36
                _ = v37
            }
            v20 = (v20 + v20)
            _ = fi29
        }
        _ = fi22
    }
    v38 := GenS1{F0: true, F1: 0}
    if (0 != -112) {
        switch "a" {
        case "hello":
            v1 = 652
        case "foo bar":
            fmt.Println(v20, v21)
        case "\n\t":
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS2{F0: 1, F1: 1000}
                var v41 uint = (v1 * 1)
                v42 := GenS1{F0: true, F1: 932}
                _ = fi39
                _ = v40
                _ = v41
                _ = v42
            }
        default:
            var v43 uint64 = 380
            _ = v43
        }
    }
    var v44 uint8 = 127
    if false {
        v21 = 2147483648
        v45 := GenS3{F0: -1, F1: ("\n\t" + "hello")}
        _ = v45
    }
    var v46 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v38
    _ = v4
    _ = v44
    _ = v46
    _ = v5
}
