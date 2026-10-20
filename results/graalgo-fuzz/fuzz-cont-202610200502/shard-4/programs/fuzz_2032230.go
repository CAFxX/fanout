package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int32 }
func genFn0(p0 int16, p1 uint) uint8 {
    v5 := GenS1{F0: -9223372036854775808}
    var v6 bool = (32768 < -1)
    v6 = (2 < 221)
    var v7 int32 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 string, p1 bool, p2 int64) string {
    switch -271 {
    case -394:
        v8 := GenS2{F0: (1 >= 924), F1: 48}
        _ = v8
    default:
        var v9 int64 = p2
        _ = v9
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 2147483647
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]int)
    v10 := GenS3{F0: ("a" + "")}
    v2 += v2
    var v11 uint8 = 255
    v11 = v11
    v2 += v2
    var v12 uint = 65535
    for fi13 := 0; fi13 < 18; fi13++ {
        var v14 uint8 = uint8(v12)
        _ = fi13
        _ = v14
    }
    var v15 bool = true
    for fi16 := 0; fi16 < 5; fi16++ {
        switch (("hello" + "a") + "a") {
        case "\n\t":
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 float64 = 0.0
                v19 := []int{(fi16 * (^(fi17))), fi16}
                if (2 < 2) {
                    fmt.Println(v18, v11, fi17)
                    var v20 int8 = 42
                    var v21 uint64 = 18446744073709551615
                    _ = v20
                    _ = v21
                } else {
                    var v22 int = -140
                    _ = v22
                }
                _ = fi17
                _ = v18
                _ = v19
            }
        case "":
            v23 := GenS2{F0: (v15 || v15), F1: 258}
            _ = v23
        default:
            var v24 uint64 = 0
            _ = v24
        }
        switch 127 {
        case -29:
            v12 += (v12 / ((v12) | 1))
        case -2147483648:
            fmt.Println(v12, v12)
        case 32768:
            if false {
                v15 = v15
                var v25 uint16 = 0
                if v15 {
                    for fi26 := 0; fi26 < 16; fi26++ {
                        v27 := GenS2{F0: true, F1: v25}
                        _ = fi26
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 14; fi28++ {
                        v29 := []int{fi28, int(v25), fi28}
                        if (v25 > v25) {
                            switch "\n\t" {
                            case "foo bar":
                                v11 = v11
                            case "\n\t":
                                v30 := GenS3{F0: "hello"}
                                _ = v30
                            }
                            v12 = 446
                            v3[uint32(v11)] = ("" + "a")
                            if ((!(!v15)) || ((true || true) && (-911 < 1))) {
                                if (!(("a" + "a") != "a")) {
                                    var v31 float64 = 1e10
                                    v15 = ((v31 * ((v31 + v31) * v31)) != v31)
                                    _ = v31
                                } else {
                                    var v32 uint8 = uint8(v12)
                                    _ = v32
                                }
                                _ = len(v29)
                            }
                        } else {
                            v2 -= (v2 + 0)
                            var v33 int = fi16
                            _ = v33
                        }
                        _ = fi28
                        _ = v29
                    }
                } else {
                    v34 := GenS1{F0: -232}
                    var v35 int8 = 100
                    v2 -= (v2 << ((845) & 63))
                    _ = v34
                    _ = v35
                }
                _ = v25
            }
        default:
            v36 := GenS4{F0: 602, F1: (v2 | (v2 / (((v2 % (((v2 + 424)) | 1))) | 1)))}
            _ = v36
        }
        _ = fi16
    }
    var v37 int64 = 9223372036854775807
    fmt.Println(v12, v11)
    switch -449 {
    case -758:
        v12 -= 983
    default:
        switch 1 {
        case -71:
            v38 := GenS1{F0: (v37 % ((v37) | 1))}
            _ = v38
        }
    }
    v12 += 0
    var v39 int16 = int16(v37)
    switch "a" {
    case "hello":
        _ = len(v1)
    case "foo bar":
        v12 = 65535
    default:
        var v40 uint16 = 0
        _ = v40
    }
    v39 += (-(v39))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v4
}
