package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64; F1 int32; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
func genFn0() uint8 {
    var v6 string = ""
    v6 = v6
    v6 = v6
    var v7 string = v6
    v8 := GenS1{F0: 2, F1: 4294967295, F2: -1}
    _ = v6
    _ = v7
    _ = v8
    return 142
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 bool = true
    var v4 float32 = 1.5
    v5 := make(map[uint32]string)
    var v9 string = ("\n\t" + ("foo bar" + ("a" + "hello")))
    if (v3 && v3) {
        if v3 {
            delete(v2, -926)
            var v10 uint = 985
            _ = v10
        }
    }
    v11 := []int{32768}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v12 int64 = 9223372036854775807
    if v3 {
        fmt.Println(v3)
        _ = len(v5)
        fmt.Println(v12, v4)
    }
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint8 = 42
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: (v1 << ((1) & 63)), F1: 0, F2: -957}
    for fi16 := 0; fi16 < 9; fi16++ {
        switch v9 {
        case "\n\t":
            delete(v5, 2)
        case "foo bar":
            v4 += (v4 * (v4 * ((((v4 * -2.25) / 1.0) + v4) - (float32(-2.25) * (1.5 / (v4 * ((v4 * v4) / ((v4 + (1.5 + 0.1)) + v4))))))))
        }
        _ = len(v11)
        _ = fi16
    }
    fmt.Println(v4)
    if true {
        if v3 {
            v17 := GenS1{F0: 548, F1: 893, F2: 42}
            if false {
                var v18 uint = 829
                var v19 uint64 = (v1 << (((v18 >> ((269) & 63))) & 63))
                v19 += (^(v19))
                _ = v18
                _ = v19
            } else {
                var v20 int64 = (v12 >> ((2) & 63))
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 uint8 = 127
                    if (849 >= uint16(v1)) {
                        v23 := GenS3{F0: (v9 + (("\n\t" + (v9 + "foo bar")) + (v9 + (((v9 + ((("hello" + v9) + "a") + ("" + (v9 + "hello")))) + (v9 + "a")) + "a"))))}
                        var v24 float32 = v4
                        fmt.Println(v22, v22, v4)
                        _ = v23
                        _ = v24
                    } else {
                        if v3 {
                            switch fi21 {
                            case 127:
                                v11 = append(v11, (fi21 & -2147483648))
                            }
                            if ((v9 + "\n\t") == v9) {
                                fmt.Println(v12, v3)
                                if v3 {
                                    var v25 float32 = (((float32(float64(v4)) + v4) / 0.1) * 0.0)
                                    if (-1 == (v20 & v12)) {
                                        v3 = false
                                    }
                                    _ = len(v11)
                                    _ = v25
                                }
                            } else {
                                var v26 bool = v3
                                _ = v26
                            }
                            v3 = v3
                        }
                    }
                    fmt.Println(v22)
                    _ = fi21
                    _ = v22
                }
                _ = v20
            }
            _ = v17
        }
        if v3 {
            if v3 {
                var v27 uint = 50
                v3 = true
                _ = v27
            } else {
                v28 := GenS4{F0: -1}
                v4 += -2.25
                _ = v28
            }
            if true {
                v3 = v3
            } else {
                var v29 string = (v9 + v9)
                _ = v29
            }
            var v30 uint64 = v1
            _ = v30
        } else {
            v4 = 1e10
            if (-9223372036854775808 < 937) {
                v31 := GenS5{F0: 2}
                var v32 uint8 = 12
                _ = v31
                _ = v32
            }
            if true {
                v9 = ((v9 + (v9 + v9)) + v9)
            } else {
                for fi33 := 0; fi33 < 10; fi33++ {
                    var v34 int64 = -109
                    _ = fi33
                    _ = v34
                }
            }
        }
        v9 = (v9 + "")
    }
    var v35 int = 9223372036854775807
    v35 += v35
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v9
}
