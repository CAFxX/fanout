package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }

func main() {
    v1 := []string{}
    var v2 uint = 115
    v3 := []int{0, 128}
    var v4 uint64 = 1
    v5 := []int{4294967295, 296, -2147483648}
    if true {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: -1}
            v2 = v2
            var v8 uint8 = 42
            _ = fi6
            _ = v7
            _ = v8
        }
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
    }
    v9 := GenS2{F0: 144}
    var v10 uint32 = 2
    fmt.Println(v10)
    v11 := GenS3{F0: 32767, F1: v4, F2: -989}
    var v12 int8 = 45
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int32 = 66
        fmt.Println(v2)
        v15 := GenS4{F0: 42, F1: (v4 + 4294967295)}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int64 = 0
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            switch (fi17 << ((255) & 63)) {
            case 1:
                var v19 int64 = -434
                _ = v19
            case 286:
                v2 = v2
            default:
                if false {
                    if false {
                        var v20 int16 = 0
                        var v21 int32 = 281
                        _ = v20
                        _ = v21
                    }
                } else {
                    v22 := GenS1{F0: v16}
                    switch -1 {
                    case 127:
                        fmt.Println(v10, fi17, v10)
                    case -2147483648:
                        var v23 uint16 = 65535
                        _ = v23
                    case 986:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        v10 = (v10 / ((2) | 1))
                    }
                    _ = v22
                }
            }
            if true {
                if false {
                    v5 = append(v5, fi18)
                    switch ("\n\t" + ("foo bar" + "a")) {
                    case "hello":
                        fmt.Println(v16, fi17, v2)
                    }
                } else {
                    if (!true) {
                        v2 += v2
                        if (true && (fi17 > (fi18 ^ 42))) {
                            fmt.Println(v2)
                            v12 -= v12
                            var v24 int = 0
                            _ = v24
                        } else {
                            var v25 uint64 = (v4 >> ((v2) & 63))
                            if len(v1) > 0 {
                                v1[0] = ""
                            }
                            fmt.Println(v2, v4, fi17)
                            _ = v25
                        }
                    } else {
                        var v26 int32 = -224
                        var v27 uint = (v2 ^ (v2 + v2))
                        v12 -= (^(v12))
                        v10 -= (v10 + 4294967295)
                        if ((((("" + "\n\t") < (("" + "hello") + "foo bar")) && true) || true) && true) {
                            v2 = v2
                        }
                        _ = v26
                        _ = v27
                    }
                    v2 -= (^(v2))
                }
                v28 := GenS2{F0: 134}
                v10 -= v10
                _ = v28
            }
            if true {
                var v29 uint32 = 2147483648
                v2 = uint(v29)
                v2 = (v2 / (((^(v2))) | 1))
                _ = v29
            }
            _ = fi18
        }
        for fi30 := 0; fi30 < 20; fi30++ {
            var v31 int8 = 1
            v32 := GenS4{F0: 104, F1: 482}
            v33 := GenS5{F0: 32767}
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        _ = fi17
    }
    if (v16 == v16) {
        v34 := GenS6{F0: 2}
        _ = len(v5)
        if ((42 >= 0) || false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v16, v2, v4)
            _ = len(v3)
        } else {
            var v35 uint = v2
            _ = v35
        }
        _ = v34
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
