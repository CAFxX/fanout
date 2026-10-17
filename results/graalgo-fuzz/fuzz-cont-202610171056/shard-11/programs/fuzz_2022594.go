package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint32 }
func genFn0() uint64 {
    var v4 bool = (!(0 != 255))
    v4 = true
    if false {
        var v5 int8 = 127
        var v6 int16 = 0
        v4 = v4
        _ = v5
        _ = v6
    } else {
        v4 = (42 < 2147483647)
        v4 = true
    }
    if (!(0 > 0)) {
        var v7 uint64 = 0
        var v8 int8 = -1
        _ = v7
        _ = v8
    }
    _ = v4
    return 760
}

func main() {
    v1 := []int{4294967295, -487}
    var v2 uint8 = 39
    v3 := make(map[int]int)
    var v9 int64 = -9223372036854775808
    var v10 int8 = -67
    var v11 uint64 = 2
    v12 := v3[174]
    v12 = v12
    fmt.Println(v11)
    v11 = genFn0()
    for fi13 := 0; fi13 < 19; fi13++ {
        switch "" {
        case "":
            switch (fi13 << ((695) & 63)) {
            case -311:
                v14 := GenS1{F0: 0, F1: (v10 << ((42) & 63)), F2: "\n\t"}
                _ = v14
            }
        }
        _ = fi13
    }
    v15 := GenS1{F0: 83, F1: int8(v9), F2: ""}
    if (true || true) {
        v16 := GenS1{F0: int32(v10), F1: (v10 ^ v10), F2: ("hello" + "hello")}
        _ = v16
    } else {
        var v17 string = (("\n\t" + ("foo bar" + "\n\t")) + "a")
        v18 := GenS1{F0: -2147483648, F1: v10, F2: ("foo bar" + "foo bar")}
        _ = v17
        _ = v18
    }
    switch (v12 & 1) {
    case 32767:
        if true {
            var v19 int32 = int32(v12)
            for fi20 := 0; fi20 < 19; fi20++ {
                for fi21 := 0; fi21 < 3; fi21++ {
                    if ((fi20 / (((fi21 << ((670) & 63))) | 1)) == v12) {
                        var v22 bool = (true && ((v11 | genFn0()) == (v11 + 960)))
                        v3[fi20] = (fi21 * fi21)
                        v23 := GenS1{F0: (v19 - (v19 / ((v19) | 1))), F1: v10, F2: "hello"}
                        _ = v22
                        _ = v23
                    }
                    var v24 uint32 = 524
                    v1 = append(v1, -880)
                    _ = fi21
                    _ = v24
                }
                switch fi20 {
                case -9223372036854775808:
                    v2 = 255
                case 0:
                    var v25 uint16 = 598
                    _ = v25
                }
                var v26 int64 = -9223372036854775808
                _ = fi20
                _ = v26
            }
            _ = v19
        } else {
            if (("" + "hello") > "") {
                var v27 string = "\n\t"
                v11 = v11
                _ = len(v3)
                _ = v27
            } else {
                if false {
                    switch (v12 / (((v12 * 4294967295)) | 1)) {
                    case 4294967295:
                        for fi28 := 0; fi28 < 6; fi28++ {
                            if (1 == 0) {
                                v29 := GenS1{F0: 1, F1: v10, F2: "\n\t"}
                                var v30 int64 = (v9 & -1)
                                v30 = (v9 >> ((4294967295) & 63))
                                _ = v29
                                _ = v30
                            } else {
                                for fi31 := 0; fi31 < 3; fi31++ {
                                    if false {
                                        _ = len(v3)
                                        v32 := GenS1{F0: -1, F1: v10, F2: ""}
                                        _ = v32
                                    }
                                    _ = len(v3)
                                    v33 := GenS1{F0: -1, F1: v10, F2: "hello"}
                                    var v34 uint64 = (^(v11))
                                    _ = fi31
                                    _ = v33
                                    _ = v34
                                }
                                for fi35 := 0; fi35 < 2; fi35++ {
                                    _ = len(v3)
                                    fmt.Println(v9, v2)
                                    _ = fi35
                                }
                                var v36 int16 = 131
                                _ = v36
                            }
                            _ = fi28
                        }
                    case -2147483648:
                        v11 = (v11 >> ((882) & 63))
                    default:
                        if ((float32(0.1) * 0.1) <= float32(3.14159)) {
                            if false {
                                v37 := GenS2{F0: int16(v9)}
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 18; fi38++ {
                                v11 = genFn0()
                                v2 += (v2 >> ((42) & 63))
                                _ = fi38
                            }
                        } else {
                            v39 := GenS1{F0: -927, F1: v10, F2: "foo bar"}
                            _ = v39
                        }
                    }
                    var v40 int64 = v9
                    _ = v40
                } else {
                    v3[v12] = (v12 | v12)
                    var v41 int8 = v10
                    _ = v41
                }
                v42 := GenS1{F0: 1, F1: v10, F2: "\n\t"}
                for fi43 := 0; fi43 < 7; fi43++ {
                    v10 -= (^(v10))
                    _ = fi43
                }
                _ = v42
            }
        }
    default:
        v11 += 18446744073709551615
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        if true {
            v2 -= v2
            var v45 int = (v12 >> ((65535) & 63))
            _ = v45
        }
        _ = fi44
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        v47 := GenS3{F0: 1}
        var v48 bool = (true && ((v10 >> ((0) & 63)) <= (^(v10))))
        _ = fi46
        _ = v47
        _ = v48
    }
    if (uint32(v11) > 617) {
        v2 = v2
        if ((1.5 != float64(float32(1.5))) || (((1.5 - float64(1e10)) < 1e10) || false)) {
            fmt.Println(v11, v11, v11)
            v3[(v12 * int(v10))] = (v12 >> ((18446744073709551615) & 63))
        } else {
            v49 := GenS4{F0: 427, F1: -1, F2: 0}
            _ = v49
        }
        _ = len(v3)
    }
    v50 := GenS1{F0: 0, F1: v10, F2: ("a" + "foo bar")}
    if false {
        for fi51 := 0; fi51 < 16; fi51++ {
            var v52 int32 = 457
            v53 := GenS1{F0: int32(v12), F1: 42, F2: "hello"}
            _ = fi51
            _ = v52
            _ = v53
        }
        switch ("\n\t" + "hello") {
        case "\n\t":
            v12 = 2147483647
        case "foo bar":
            v2 -= (v2 * (v2 & (v2 ^ (^(v2)))))
        }
        var v54 uint16 = 849
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v50
    _ = v9
}
