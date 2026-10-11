package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16; F1 int; F2 int64 }
type GenS8 struct { F0 uint64; F1 string; F2 bool }

func main() {
    var v1 int16 = -350
    var v2 uint32 = 42
    var v3 float32 = 1e10
    var v4 uint32 = v2
    v5 := make(map[string]string)
    v6 := make(map[uint32]int)
    if true {
        var v7 uint16 = 1
        v8 := GenS1{F0: 35, F1: 2147483647, F2: 1}
        fmt.Println(v3, v3)
        _ = v7
        _ = v8
    } else {
        var v9 int16 = (v1 / ((v1) | 1))
        _ = v9
    }
    v10 := GenS1{F0: 4294967295, F1: 1, F2: -861}
    if ((v4 & v2) > v2) {
        if true {
            if (true || true) {
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 int = 32768
                    v1 -= v1
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 4; fi13++ {
                    v14 := GenS2{F0: "\n\t", F1: 0}
                    fmt.Println(v3, fi13)
                    var v15 uint = 1
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: 179, F1: 127, F2: -9223372036854775808}
                _ = v16
            } else {
                if false {
                    var v17 uint16 = 42
                    v1 = 0
                    _ = v17
                }
            }
            var v18 int64 = int64(v4)
            _ = v18
        } else {
            if (-43 >= -687) {
                v3 -= v3
                v4 += v4
                v19 := GenS2{F0: "hello", F1: -9223372036854775808}
                _ = v19
            } else {
                v5["\n\t"] = "a"
                var v20 uint8 = 2
                v21 := GenS3{F0: 127, F1: v20, F2: (v4 + v2)}
                _ = v20
                _ = v21
            }
            if (65535 > 0) {
                var v22 int64 = int64(v1)
                v23 := GenS3{F0: 32767, F1: 127, F2: 65535}
                if false {
                    var v24 uint16 = 0
                    _ = v24
                } else {
                    var v25 uint64 = 0
                    if false {
                        var v26 float32 = float32(1.5)
                        if true {
                            v26 -= (v26 + (v26 + v26))
                        }
                        _ = v26
                    }
                    _ = v25
                }
                _ = v22
                _ = v23
            }
            v27 := GenS3{F0: 32768, F1: 128, F2: v4}
            _ = v27
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            if (!true) {
                v4 = v2
                v3 = 1.5
                switch (("\n\t" + (("hello" + "a") + "a")) + (("hello" + "\n\t") + "foo bar")) {
                case "hello":
                    var v29 int8 = 0
                    _ = v29
                case "":
                    switch "" {
                    case "":
                        v3 -= (v3 * (v3 - v3))
                    case "\n\t":
                        if false {
                            var v30 int16 = v1
                            v31 := v5[("hello" + (("foo bar" + "") + ("foo bar" + ("a" + ("a" + "hello")))))]
                            _ = v30
                            _ = v31
                        }
                    case "hello":
                        var v32 uint64 = 4294967295
                        _ = v32
                    }
                case "a":
                    if false {
                        if false {
                            fmt.Println(v4, v3, fi28)
                            if (!(false || (!true))) {
                                switch fi28 {
                                case -2147483648:
                                    fmt.Println(v3, fi28, fi28)
                                case 0:
                                    v3 += v3
                                }
                                var v33 uint = 558
                                v34 := GenS4{F0: uint32(v33)}
                                _ = v33
                                _ = v34
                            } else {
                                v35 := GenS5{F0: 239}
                                switch fi28 {
                                case 4294967295:
                                    var v36 uint = 0
                                    _ = v36
                                }
                                for fi37 := 0; fi37 < 1; fi37++ {
                                    var v38 uint16 = 0
                                    v3 -= v3
                                    _ = fi37
                                    _ = v38
                                }
                                _ = v35
                            }
                            var v39 uint64 = 0
                            _ = v39
                        } else {
                            v40 := []int{fi28}
                            _ = v40
                        }
                        v41 := []int{32767}
                        v2 = 2
                        _ = v41
                    }
                default:
                    var v42 int16 = (-(v1))
                    _ = v42
                }
            }
            _ = fi28
        }
        v3 += v3
    }
    v43 := GenS2{F0: "\n\t", F1: 330}
    v44 := GenS6{F0: -679}
    var v45 int = 2147483647
    var v46 int16 = 849
    v6[(v2 & (v4 % ((v4) | 1)))] = (v45 << ((65535) & 63))
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 string = "a"
        v49 := GenS3{F0: fi47, F1: 1, F2: (v2 / ((0) | 1))}
        _ = fi47
        _ = v48
        _ = v49
    }
    if true {
        switch "hello" {
        case "hello":
            v2 = (v2 + (v2 >> ((177) & 63)))
        default:
            var v50 int32 = -1
            _ = v50
        }
        for fi51 := 0; fi51 < 13; fi51++ {
            if false {
                var v52 uint32 = (v4 | (v2 + (^(v4))))
                _ = v52
            }
            _ = fi51
        }
    }
    v4 += (v4 | 968)
    v53 := []int{v45, v45, (v45 & v45)}
    v1 += (v1 | (v1 * v1))
    if (true || (!true)) {
        v54 := GenS7{F0: (v1 - v46), F1: -231, F2: -9223372036854775808}
        v4 = v4
        for fi55 := 0; fi55 < 9; fi55++ {
            for fi56 := 0; fi56 < 3; fi56++ {
                v57 := GenS2{F0: ("foo bar" + "hello"), F1: 9223372036854775807}
                v45 -= 758
                v58 := GenS8{F0: 42, F1: "foo bar", F2: false}
                var v59 bool = (v46 != v46)
                v59 = ((!(255 <= 42)) && (543 < -2147483648))
                fmt.Println(v59, fi55)
                var v60 int = (v45 + (fi55 ^ (fi55 >> ((284) & 63))))
                _ = fi56
                _ = v57
                _ = v58
                _ = v59
                _ = v60
            }
            v46 = v1
            v4 += v2
            _ = fi55
        }
        _ = v54
    }
    v4 += (v4 & 1)
    var v61 int8 = 100
    var v62 uint = 65535
    if len(v53) > 0 {
        v53[0] = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v53
    _ = v6
    _ = v61
    _ = v62
}
