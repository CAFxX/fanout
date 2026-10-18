package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int8; F1 uint; F2 uint32 }
type GenS4 struct { F0 int; F1 string }
func genFn0(p0 uint64, p1 string) string {
    fmt.Println(p0, p0)
    p0 = 4294967295
    v7 := []int{42, 906}
    _ = p0
    _ = p1
    _ = v7
    return (p1 + ((p1 + (("" + "a") + "foo bar")) + p1))
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[string]int)
    var v3 float32 = 3.14159
    var v4 bool = (424 >= -1)
    v5 := []uint32{65535, 65535, 2, 65535}
    var v6 int = -960
    var v8 int64 = -876
    v9 := GenS1{F0: 128, F1: 0, F2: v8}
    fmt.Println(v6, v4, v8)
    for fi10 := 0; fi10 < 11; fi10++ {
        switch "" {
        case "hello":
            v11 := GenS1{F0: 1, F1: 66, F2: -1}
            _ = v11
        case "\n\t":
            v2["\n\t"] = (fi10 ^ (fi10 * -9223372036854775808))
        case "foo bar":
            if (18446744073709551615 > 349) {
                var v12 uint32 = uint32(v1)
                _ = v12
            }
        }
        _ = fi10
    }
    v4 = (((69 >= 127) || false) && (167 <= 1))
    v1 -= (v1 % (((v1 / ((255) | 1))) | 1))
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int16 = -32768
        if false {
            v4 = (!(v8 < (v8 - v8)))
            var v15 uint64 = 2
            v5 = append(v5, uint32(v1))
            _ = v15
        } else {
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
            var v16 int16 = (v14 ^ (-(v14)))
            _ = v16
        }
        switch v6 {
        case 0:
            var v17 uint32 = 1
            _ = v17
        case -2147483648:
            if v4 {
                v4 = (255 != 139)
                v18 := GenS1{F0: 171, F1: 373, F2: (v8 >> ((0) & 63))}
                if (!true) {
                    var v19 int32 = -804
                    _ = v19
                } else {
                    var v20 uint8 = 12
                    if (((v1 << ((4294967295) & 63)) > (v1 + (v1 + v1))) && false) {
                        for fi21 := 0; fi21 < 19; fi21++ {
                            v1 += (v1 - 1)
                            switch (("\n\t" + "") + "a") {
                            case "":
                                var v22 int8 = -105
                                _ = v22
                            }
                            var v23 bool = v4
                            _ = fi21
                            _ = v23
                        }
                        v3 += 0.1
                        if (2 >= 292) {
                            v8 -= v8
                            v1 += v1
                        }
                    } else {
                        v24 := GenS1{F0: v20, F1: 197, F2: v8}
                        if v4 {
                            for fi25 := 0; fi25 < 10; fi25++ {
                                v26 := GenS1{F0: (v20 | v20), F1: 65535, F2: v8}
                                switch ("" + ("hello" + ("a" + "\n\t"))) {
                                case "hello":
                                    if v4 {
                                        if true {
                                            v5 = append(v5, 2)
                                        }
                                        fmt.Println(v8, v1)
                                        if v4 {
                                            var v27 int32 = 42
                                            switch ("foo bar" + ("" + ("" + "foo bar"))) {
                                            case "a":
                                                if v4 {
                                                    _ = len(v5)
                                                }
                                            }
                                            if (42 != v1) {
                                                var v28 uint16 = v1
                                                switch (fi25 / ((fi13) | 1)) {
                                                case 9223372036854775807:
                                                    fmt.Println(v1)
                                                case 1:
                                                    var v29 uint64 = 656
                                                    _ = v29
                                                case 32768:
                                                    v30 := GenS2{F0: (^(fi25))}
                                                    _ = v30
                                                }
                                                _ = v28
                                            }
                                            _ = v27
                                        } else {
                                            var v31 uint8 = 127
                                            v31 += (v20 + v20)
                                            if ((v8 / ((42) | 1)) > (v8 + -688)) {
                                                switch (fi25 + 4294967295) {
                                                case 32767:
                                                    var v32 float64 = 0.0
                                                    _ = v32
                                                default:
                                                    fmt.Println(fi25, v14, v4)
                                                }
                                            }
                                            _ = v31
                                        }
                                    }
                                case "\n\t":
                                    var v33 uint8 = (v20 << ((2) & 63))
                                    _ = v33
                                case "foo bar":
                                    _ = len(v2)
                                    v34 := GenS2{F0: fi13}
                                    _ = v34
                                default:
                                    v35 := GenS2{F0: 0}
                                    _ = v35
                                }
                                _ = fi25
                                _ = v26
                            }
                            var v36 int32 = 42
                            _ = v36
                        } else {
                            v8 -= v8
                        }
                        _ = v24
                    }
                    _ = v20
                }
                _ = v18
            } else {
                v8 = (^(v8))
                v37 := GenS3{F0: -1, F1: 0, F2: 1}
                _ = v37
            }
        }
        _ = fi13
        _ = v14
    }
    _ = len(v2)
    v38 := GenS4{F0: v6, F1: "\n\t"}
    _ = len(v2)
    v39 := GenS1{F0: 42, F1: 0, F2: 441}
    v40 := v2[""]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
