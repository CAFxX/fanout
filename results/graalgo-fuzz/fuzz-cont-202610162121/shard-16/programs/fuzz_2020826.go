package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8) bool {
    if false {
        v7 := []int{-2147483648, -590, 4294967295, -604}
        var v8 bool = true
        if (100 > -128) {
            v9 := GenS1{F0: 29}
            var v10 int32 = 1
            v10 += v10
            _ = v10
            _ = v9
        }
        _ = v7
        _ = v8
    }
    p0 = (p0 / (((p0 / (((p0 + p0)) | 1))) | 1))
    _ = p0
    return true
}

func main() {
    var v1 string = "a"
    var v2 int8 = -128
    var v3 int32 = -1
    v4 := make(map[uint32]int)
    var v5 uint32 = 488
    v6 := make(map[string]int)
    var v11 uint8 = 103
    v12 := []int{-1}
    var v13 int = 4294967295
    var v14 int = int(v3)
    _ = len(v12)
    if false {
        for fi15 := 0; fi15 < 3; fi15++ {
            v2 += v2
            switch v13 {
            case 9223372036854775807:
                for fi16 := 0; fi16 < 17; fi16++ {
                    var v17 int32 = v3
                    _ = fi16
                    _ = v17
                }
            case 833:
                var v18 bool = (!(1e10 < 0.1))
                _ = v18
            }
            for fi19 := 0; fi19 < 15; fi19++ {
                if false {
                    var v20 float32 = float32(0.0)
                    if (!(-1 <= (v2 & 0))) {
                        if false {
                            var v21 int32 = (v3 >> ((2) & 63))
                            switch (v1 + (v1 + (v1 + (v1 + v1)))) {
                            case "":
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                            v20 = 1e10
                            _ = v21
                        } else {
                            v3 += (-(v3))
                        }
                    }
                    _ = v20
                }
                v5 += 42
                switch (v14 & v14) {
                case 1:
                    v14 = (v14 << ((161) & 63))
                default:
                    var v22 float32 = (3.14159 + 0.1)
                    _ = v22
                }
                _ = fi19
            }
            _ = fi15
        }
    }
    if (!(2 > uint64(v5))) {
        if true {
            _ = len(v4)
            switch v1 {
            case "\n\t":
                if false {
                    v14 -= (v14 % ((v14) | 1))
                }
            }
        }
        delete(v4, (v5 + 4294967295))
        for fi23 := 0; fi23 < 18; fi23++ {
            switch v1 {
            case "\n\t":
                v3 = (v3 ^ (v3 % ((-157) | 1)))
            }
            if true {
                v24 := GenS1{F0: v2}
                if true {
                    v3 += -2147483648
                    delete(v4, 65535)
                }
                _ = v24
            } else {
                var v25 uint = 42
                _ = v25
            }
            var v26 int32 = v3
            _ = fi23
            _ = v26
        }
    }
    if true {
        v13 -= (v14 ^ (v13 % ((v13) | 1)))
    } else {
        v3 -= (v3 ^ -1)
        if (!false) {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        } else {
            v11 = 127
            var v27 int16 = 42
            v13 += v14
            _ = v27
        }
        if (0.1 == 3.14159) {
            var v28 int = -9223372036854775808
            _ = v28
        }
    }
    if true {
        v11 = 128
        if (1e10 > ((1e10 / (1e10 + 0.1)) - 0.0)) {
            v29 := GenS2{F0: v11}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                v1 = (v1 + "a")
                _ = len(v4)
                _ = fi30
            }
            switch v1 {
            case "":
                for fi31 := 0; fi31 < 8; fi31++ {
                    fmt.Println(v2)
                    _ = fi31
                }
            case "foo bar":
                if (!true) {
                    fmt.Println(v2, v11)
                    var v32 int32 = (v3 + 339)
                    _ = v32
                }
            }
        }
    }
    v33 := v4[v5]
    var v34 int64 = int64(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
