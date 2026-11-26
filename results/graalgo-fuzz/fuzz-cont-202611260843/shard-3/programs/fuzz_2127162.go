package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint16; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 int16; F2 int32 }

func main() {
    var v1 float64 = (0.0 + 1.5)
    v2 := make(map[string]string)
    var v3 int64 = 1
    v4 := make(map[int]bool)
    v5 := []uint32{742}
    var v6 uint = 42
    v7 := v2["hello"]
    fmt.Println(v7)
    if (206 != 1) {
        v8 := GenS1{F0: -128}
        _ = v8
    }
    if false {
        var v9 bool = (127 >= 255)
        _ = v9
    }
    switch 32767 {
    case -996:
        v10 := GenS1{F0: 100}
        _ = v10
    default:
        var v11 uint = 18446744073709551615
        _ = v11
    }
    switch v7 {
    case "a":
        var v12 uint8 = uint8(v6)
        _ = v12
    case "hello":
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                if (2 >= 1) {
                    var v15 bool = (!true)
                    if v15 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        var v16 string = ""
                        _ = v16
                    } else {
                        var v17 uint64 = 968
                        _ = v17
                    }
                    _ = v15
                } else {
                    if (false && false) {
                        var v18 int32 = -89
                        if true {
                            if true {
                                v18 = v18
                                var v19 uint32 = 2147483647
                                var v20 uint32 = v19
                                _ = v19
                                _ = v20
                            } else {
                                if (1 < 4294967295) {
                                    v6 = (v6 << (((^(v6))) & 63))
                                }
                            }
                        } else {
                            if true {
                                var v21 int16 = -455
                                var v22 int16 = -1
                                if (127 < -35) {
                                    switch fi14 {
                                    case 341:
                                        var v23 int16 = v21
                                        _ = v23
                                    default:
                                        v24 := GenS2{F0: false, F1: 0, F2: (v22 - 32767)}
                                        _ = v24
                                    }
                                }
                                var v25 float32 = 0.1
                                v26 := GenS3{F0: -1}
                                _ = v21
                                _ = v22
                                _ = v25
                                _ = v26
                            }
                            if (true || true) {
                                var v27 int = int(v3)
                                _ = v27
                            } else {
                                var v28 int16 = -32768
                                _ = v28
                            }
                            var v29 int = 743
                            _ = v29
                        }
                        _ = v18
                    } else {
                        var v30 int32 = -172
                        if (0 == 2) {
                            var v31 float64 = v1
                            v32 := GenS4{F0: (v30 ^ 42), F1: 42, F2: v30}
                            _ = v31
                            _ = v32
                        } else {
                            v6 = 255
                            if (true && (65535 != v6)) {
                                var v33 uint32 = 2
                                if ((!true) || true) {
                                    var v34 uint = 914
                                    v34 -= (^(v6))
                                    _ = v34
                                }
                                delete(v4, 4294967295)
                                _ = v33
                            } else {
                                var v35 uint32 = 231
                                v30 -= v30
                                v35 = (v35 - 1)
                                _ = v35
                            }
                        }
                        delete(v4, fi13)
                        _ = v30
                    }
                    var v36 uint32 = 172
                    v5 = append(v5, v36)
                    _ = v36
                }
                var v37 int16 = -1
                _ = fi14
                _ = v37
            }
            _ = fi13
        }
    case "foo bar":
        var v38 int16 = -32768
        _ = v38
    default:
        v6 = (v6 | v6)
    }
    switch 801 {
    case -9223372036854775808:
        v1 += 1.5
    case 42:
        v6 -= (^(v6))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
