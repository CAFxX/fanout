package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint; F2 uint }

func main() {
    var v1 int64 = 0
    v2 := []uint32{522, 2, 4294967295}
    var v3 uint8 = 34
    var v4 int = 42
    v5 := make(map[string]bool)
    fmt.Println(v3, v1)
    var v6 uint64 = 42
    v7 := GenS1{F0: (v1 - v1)}
    var v8 uint32 = 2
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int = (v4 & (v4 >> ((67) & 63)))
        var v11 float32 = 1.5
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            if true {
                var v14 uint8 = v3
                v4 += fi13
                var v15 int32 = -1
                _ = v14
                _ = v15
            } else {
                fmt.Println(v8, v8)
                var v16 float64 = (-2.25 + (0.1 * 0.0))
                _ = v16
            }
            var v17 uint8 = (v3 << ((701) & 63))
            v18 := GenS1{F0: -588}
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi13
            _ = v17
            _ = v18
        }
        var v19 int8 = 42
        _ = fi12
        _ = v19
    }
    if false {
        var v20 int16 = 1000
        _ = v20
    } else {
        if false {
            fmt.Println(v4)
        }
        var v21 uint8 = v3
        _ = v21
    }
    v4 = v4
    if true {
        var v22 uint = uint(v8)
        if false {
            v22 = 1
            fmt.Println(v4, v6)
        } else {
            var v23 bool = (true || true)
            _ = v23
        }
        if false {
            v24 := GenS1{F0: v1}
            var v25 int32 = 42
            if false {
                for fi26 := 0; fi26 < 10; fi26++ {
                    v27 := GenS2{F0: 0}
                    _ = fi26
                    _ = v27
                }
            } else {
                var v28 uint8 = (v3 ^ (v3 >> ((v22) & 63)))
                delete(v5, (("\n\t" + "a") + "\n\t"))
                v8 = 40
                _ = v28
            }
            _ = v24
            _ = v25
        } else {
            v29 := GenS1{F0: v1}
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 int = (-(fi30))
                v32 := GenS1{F0: (v1 + int64(v6))}
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = v29
        }
        _ = v22
    } else {
        var v33 uint16 = 1
        for fi34 := 0; fi34 < 20; fi34++ {
            if (((1e10 + 3.14159) / 0.1) == float32(3.14159)) {
                v35 := GenS3{F0: uint(v3), F1: 255, F2: 694}
                if false {
                    v2 = append(v2, (v8 - v8))
                } else {
                    _ = len(v5)
                    if false {
                        var v36 int = 42
                        _ = v36
                    }
                    var v37 int32 = 0
                    _ = v37
                }
                _ = v35
            }
            if false {
                v6 = v6
                if (32767 < -32768) {
                    v38 := GenS3{F0: 909, F1: 42, F2: 346}
                    v39 := GenS1{F0: -9223372036854775808}
                    if false {
                        v6 = v6
                        var v40 float64 = (0.0 + (float64(((1.5 / 1.0) / 1.0)) - 1e10))
                        _ = v40
                    } else {
                        v33 += 2
                        switch -1 {
                        case 42:
                            for fi41 := 0; fi41 < 10; fi41++ {
                                v2 = append(v2, (v8 + v8))
                                if true {
                                    v1 -= v1
                                    var v42 uint8 = (v3 - (v3 ^ 2))
                                    _ = v42
                                } else {
                                    v6 -= (v6 + (^(v6)))
                                }
                                var v43 uint32 = (v8 - v8)
                                _ = fi41
                                _ = v43
                            }
                        default:
                            v44 := GenS1{F0: (v1 + (v1 | v1))}
                            _ = v44
                        }
                        v4 += (fi34 - (fi34 | v4))
                    }
                    _ = v38
                    _ = v39
                }
            } else {
                switch "a" {
                case "hello":
                    v3 = v3
                }
                var v45 float64 = ((1e10 / 0.1) - ((0.1 / float64((3.14159 - (1.5 - 0.0)))) * 0.0))
                v45 = v45
                _ = v45
            }
            v46 := GenS1{F0: (v1 << ((0) & 63))}
            _ = fi34
            _ = v46
        }
        for fi47 := 0; fi47 < 6; fi47++ {
            v48 := v5["hello"]
            switch "hello" {
            case "a":
                switch "" {
                case "hello":
                    var v49 uint64 = v6
                    _ = v49
                case "a":
                    var v50 string = "hello"
                    _ = v50
                case "foo bar":
                    v51 := GenS1{F0: -1}
                    _ = v51
                }
            }
            _ = fi47
            _ = v48
        }
        _ = v33
    }
    var v52 float64 = 1e10
    var v53 string = "a"
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v7
    _ = v8
}
