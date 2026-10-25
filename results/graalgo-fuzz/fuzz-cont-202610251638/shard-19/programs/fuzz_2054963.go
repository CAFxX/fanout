package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 bool, p1 int32, p2 bool) int32 {
    v5 := GenS1{F0: 0}
    p1 += (p1 + p1)
    var v6 int = 637
    v7 := []int{(v6 / (((^(v6))) | 1)), (v6 / (((^(v6))) | 1)), (-(v6))}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 int, p1 uint8) string {
    p0 -= p0
    if true {
        if (1 > 32767) {
            p0 = 42
            if (1 <= 504) {
                v8 := []int{2147483647}
                if ((0 != 255) && ("\n\t" == "hello")) {
                    v8 = append(v8, p0)
                }
                _ = v8
            } else {
                var v9 uint8 = 188
                for fi10 := 0; fi10 < 9; fi10++ {
                    if false {
                        v11 := GenS2{F0: -57, F1: genFn0((42 >= 4294967295), 0, (255 <= 386))}
                        if ((fi10 / ((p0) | 1)) < (^(fi10))) {
                            v12 := GenS3{F0: (^(v9)), F1: (!true), F2: 18446744073709551615}
                            _ = v12
                        }
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v14 := []int{int(v9), p0, p0, (fi10 - p0)}
                            p0 = p0
                            p1 += (v9 | p1)
                            _ = fi13
                            _ = v14
                        }
                        _ = v11
                    }
                    _ = fi10
                }
                _ = v9
            }
        } else {
            v15 := GenS2{F0: 42, F1: 0}
            _ = v15
        }
        p1 -= 0
    }
    v16 := GenS1{F0: 42}
    v17 := GenS3{F0: (p1 / ((0) | 1)), F1: (true && false), F2: 89}
    v18 := GenS4{F0: ("\n\t" + ""), F1: -32768}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return "\n\t"
}

func main() {
    var v1 int = -61
    var v2 uint32 = 1
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]int)
    fmt.Println(v3, v3)
    for fi19 := 0; fi19 < 2; fi19++ {
        var v20 uint32 = v2
        v21 := []int{fi19, -651, (fi19 >> ((uint(v2)) & 63))}
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 uint64 = (v3 ^ (^(v3)))
        _ = fi22
        _ = v23
    }
    var v24 uint16 = 65535
    v25 := GenS5{F0: 2}
    fmt.Println(v2, v2, v3)
    for fi26 := 0; fi26 < 4; fi26++ {
        var v27 int64 = -1
        var v28 int64 = (v27 >> ((4294967295) & 63))
        if true {
            v2 = (v2 | v2)
        } else {
            for fi29 := 0; fi29 < 18; fi29++ {
                switch "a" {
                case "hello":
                    if ((v2 / ((551) | 1)) <= v2) {
                        v24 = v24
                        fmt.Println(fi26, v28)
                    } else {
                        var v30 int64 = (v28 / (((v27 ^ v28)) | 1))
                        v31 := []int{(^(fi26)), 1, (^(fi29)), fi29}
                        _ = v30
                        _ = v31
                    }
                case "\n\t":
                    if ((0 == -820) || (!false)) {
                        fmt.Println(fi26)
                        if ("" >= "a") {
                            v32 := GenS1{F0: v27}
                            _ = len(v4)
                            _ = v32
                        } else {
                            switch 728 {
                            case -188:
                                var v33 uint64 = (^(v3))
                                v33 = v33
                                _ = v33
                            case 128:
                                v27 = v28
                            }
                            var v34 uint8 = uint8(v3)
                            _ = v34
                        }
                    }
                case "foo bar":
                    v2 += v2
                }
                var v35 uint = 65535
                _ = fi29
                _ = v35
            }
            v2 = v2
            v28 = (v28 % ((v27) | 1))
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    switch (^(v1)) {
    case 4294967295:
        switch (("foo bar" + ("foo bar" + "foo bar")) + ("foo bar" + "\n\t")) {
        case "foo bar":
            if false {
                if false {
                    var v36 uint8 = 42
                    _ = v36
                }
                var v37 uint8 = 255
                var v38 uint64 = 0
                _ = v37
                _ = v38
            } else {
                _ = len(v4)
            }
        }
    case 41:
        v24 = (v24 ^ (^(v24)))
    }
    var v39 bool = false
    var v40 float32 = 0.0
    v41 := []int{(v1 | v1), 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
