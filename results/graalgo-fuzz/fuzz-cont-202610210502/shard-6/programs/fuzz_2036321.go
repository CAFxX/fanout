package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int16 }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0() uint8 {
    v7 := GenS1{F0: 100, F1: -1, F2: 2147483648}
    if true {
        v8 := []int{-68}
        var v9 uint = 42
        var v10 uint32 = 277
        _ = v10
        _ = v8
        _ = v9
    } else {
        if true {
            var v11 bool = (0 >= 0)
            v12 := GenS1{F0: -1, F1: 32768, F2: 65535}
            switch "hello" {
            case "a":
                for fi13 := 0; fi13 < 11; fi13++ {
                    v11 = v11
                    _ = fi13
                }
            }
            _ = v11
            _ = v12
        }
    }
    _ = v7
    return 42
}
func genFn1(p0 uint16) bool {
    p0 -= p0
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int32 = -1
        p0 -= (p0 >> ((321) & 63))
        if false {
            if false {
                var v16 uint8 = 1
                v17 := []int{}
                v17 = append(v17, fi14)
                _ = v16
                _ = v17
            } else {
                var v18 uint64 = 42
                v18 -= v18
                _ = v18
            }
            p0 += p0
        } else {
            var v19 int8 = -97
            p0 = 32767
            _ = v19
        }
        _ = fi14
        _ = v15
    }
    var v20 uint8 = 121
    for fi21 := 0; fi21 < 8; fi21++ {
        v20 -= v20
        p0 -= uint16(v20)
        v20 = (^(v20))
        _ = fi21
    }
    _ = p0
    _ = v20
    return false
}

func main() {
    var v1 int16 = -81
    v2 := make(map[string]string)
    var v3 int16 = v1
    var v4 uint16 = 255
    v5 := make(map[uint32]string)
    var v6 uint16 = (^(v4))
    var v22 bool = true
    v23 := []int{42, 4294967295, 42, 1}
    fmt.Println(v1, v4)
    v3 += 494
    v24 := GenS1{F0: -1, F1: 32768, F2: 53}
    var v25 int16 = 285
    fmt.Println(v3, v6, v25)
    var v26 int32 = 42
    v26 -= v26
    for fi27 := 0; fi27 < 10; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            if len(v23) > 0 {
                v23[0] = -727
            }
            _ = fi28
        }
        var v29 uint16 = (^(v6))
        delete(v5, 2147483647)
        _ = fi27
        _ = v29
    }
    if ((false && v22) || ("foo bar" < "")) {
        v30 := GenS1{F0: 127, F1: 32767, F2: 2}
        var v31 bool = (v22 && (false || true))
        var v32 string = "\n\t"
        _ = v30
        _ = v31
        _ = v32
    }
    if (0.1 <= 0.1) {
        v2["foo bar"] = "a"
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            if v22 {
                if (1 >= 134) {
                    fmt.Println(v6, v26)
                }
                for fi34 := 0; fi34 < 11; fi34++ {
                    switch fi33 {
                    case 4294967295:
                        v2["foo bar"] = "hello"
                    }
                    _ = len(v5)
                    v35 := GenS1{F0: 0, F1: fi34, F2: 1}
                    _ = fi34
                    _ = v35
                }
                var v36 int16 = v1
                _ = v36
            } else {
                if len(v23) > 0 {
                    v23[0] = (fi33 * 32768)
                }
                if (1 != 2) {
                    var v37 uint64 = 2
                    _ = v37
                } else {
                    v38 := GenS2{F0: -14, F1: 32767, F2: 32767}
                    switch (fi33 & (fi33 - (fi33 >> ((4294967295) & 63)))) {
                    case 0:
                        var v39 uint32 = 4294967295
                        _ = v39
                    case 128:
                        switch "" {
                        case "foo bar":
                            if len(v23) > 0 {
                                v23[0] = (fi33 - int(v26))
                            }
                        case "":
                            v40 := v5[2]
                            _ = v40
                        default:
                            for fi41 := 0; fi41 < 12; fi41++ {
                                _ = len(v2)
                                _ = len(v23)
                                v42 := GenS1{F0: 100, F1: fi41, F2: 4294967295}
                                _ = fi41
                                _ = v42
                            }
                        }
                    case 2147483647:
                        v43 := GenS1{F0: 0, F1: fi33, F2: 42}
                        _ = v43
                    default:
                        for fi44 := 0; fi44 < 20; fi44++ {
                            _ = len(v2)
                            _ = fi44
                        }
                    }
                    _ = v38
                }
            }
            for fi45 := 0; fi45 < 2; fi45++ {
                var v46 string = "a"
                if (((0.0 + 3.14159) + 0.0) >= 1.5) {
                    fmt.Println(v3, v25)
                    fmt.Println(v1, v4)
                    switch (v46 + ("hello" + v46)) {
                    case "a":
                        fmt.Println(v1, v26)
                    }
                } else {
                    var v47 uint32 = 21
                    fmt.Println(v25)
                    _ = v47
                }
                switch int(v4) {
                case -938:
                    if v22 {
                        var v48 bool = (((-(v25)) <= v25) && true)
                        v49 := GenS3{F0: v22, F1: v4}
                        var v50 bool = v22
                        _ = v48
                        _ = v49
                        _ = v50
                    } else {
                        var v51 uint16 = 65535
                        v52 := GenS3{F0: false, F1: 57}
                        _ = v51
                        _ = v52
                    }
                }
                _ = fi45
                _ = v46
            }
            _ = fi33
        }
    }
    v25 -= (v25 ^ -1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
