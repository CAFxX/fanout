package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool }
func genFn0() int8 {
    v4 := []int{891}
    if (65535 > 283) {
        var v5 float32 = (1.5 + 1.5)
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 float64 = 0.0
            _ = fi6
            _ = v7
        }
        if len(v4) > 0 {
            v4[0] = -1
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 2; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (-2.25 < 0.0) {
                for fi9 := 0; fi9 < 10; fi9++ {
                    if len(v4) > 0 {
                        v4[0] = -232
                    }
                    var v10 uint32 = 630
                    if len(v4) > 0 {
                        v4[0] = (fi9 << ((42) & 63))
                    }
                    v10 = (v10 - 2)
                    var v11 uint16 = 65535
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                var v12 int = 0
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 int32 = -153
                    var v15 uint16 = 65535
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                v16 := GenS1{F0: fi8}
                v17 := GenS1{F0: fi8}
                _ = v16
                _ = v17
            }
            _ = fi8
        }
        for fi18 := 0; fi18 < 18; fi18++ {
            if (!(-81 >= -1)) {
                v19 := GenS2{F0: -105, F1: (2 <= 127)}
                switch (("foo bar" + ("a" + "\n\t")) + "hello") {
                case "":
                    var v20 int64 = 42
                    _ = v20
                case "a":
                    var v21 uint16 = 255
                    _ = v21
                default:
                    var v22 bool = (-1 < 1000)
                    _ = v22
                }
                _ = v19
            } else {
                for fi23 := 0; fi23 < 1; fi23++ {
                    if len(v4) > 0 {
                        v4[0] = (^(fi18))
                    }
                    var v24 int = 0
                    _ = len(v4)
                    _ = fi23
                    _ = v24
                }
                var v25 float64 = float64(0.1)
                _ = v25
            }
            var v26 int64 = 263
            var v27 int64 = 422
            _ = fi18
            _ = v26
            _ = v27
        }
    }
    var v28 uint32 = 2147483648
    _ = v28
    _ = v4
    return 36
}

func main() {
    var v1 int32 = -1
    var v2 uint64 = 0
    var v3 int8 = 71
    var v29 string = "foo bar"
    var v30 uint16 = 32767
    var v31 int16 = 32767
    fmt.Println(v31, v29, v1)
    var v32 uint32 = 4294967295
    var v33 uint64 = v2
    if true {
        v34 := GenS3{F0: true}
        _ = v34
    } else {
        var v35 uint16 = 0
        for fi36 := 0; fi36 < 5; fi36++ {
            var v37 int = fi36
            _ = fi36
            _ = v37
        }
        fmt.Println(v29, v1)
        _ = v35
    }
    var v38 float32 = 0.1
    var v39 uint64 = (v33 >> ((593) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v39
}
