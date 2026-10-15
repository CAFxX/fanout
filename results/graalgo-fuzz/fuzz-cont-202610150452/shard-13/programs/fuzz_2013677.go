package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) uint64 {
    v6 := []int{1, 32768, -1}
    var v7 int64 = 9223372036854775807
    var v8 string = ""
    fmt.Println(p0, v8)
    var v9 uint64 = p0
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 int32 = -489
    var v4 float64 = 0.0
    var v5 uint16 = 5
    var v10 int32 = (v3 / (((v3 % (((v3 - 42)) | 1))) | 1))
    v3 -= 1
    var v11 bool = false
    if (!false) {
        var v12 uint8 = 42
        _ = v12
    }
    if ((-94 > 100) || false) {
        for fi13 := 0; fi13 < 7; fi13++ {
            fmt.Println(v11)
            v14 := GenS1{F0: 42}
            v15 := v1[2147483648]
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v16 := []int{}
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 12; fi18++ {
            delete(v2, 1)
            if (v4 >= float64((1.5 - -2.25))) {
                if true {
                    var v19 int64 = 9223372036854775807
                    var v20 float32 = -2.25
                    _ = v19
                    _ = v20
                }
                var v21 float64 = (-2.25 - 0.0)
                _ = v21
            }
            var v22 float32 = ((3.14159 - 3.14159) * 3.14159)
            _ = fi18
            _ = v22
        }
        var v23 int16 = 32767
        var v24 string = ("" + "")
        _ = fi17
        _ = v23
        _ = v24
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        _ = len(v16)
        switch "" {
        case "\n\t":
            var v26 uint32 = 65535
            _ = v26
        case "hello":
            for fi27 := 0; fi27 < 14; fi27++ {
                var v28 uint16 = v5
                var v29 int32 = v10
                _ = len(v16)
                _ = fi27
                _ = v28
                _ = v29
            }
        case "foo bar":
            switch (fi25 - fi25) {
            case -9223372036854775808:
                for fi30 := 0; fi30 < 16; fi30++ {
                    _ = len(v16)
                    fmt.Println(v3)
                    if (106 >= 127) {
                        fmt.Println(v3)
                        v5 = 301
                        fmt.Println(v10, v4, fi30)
                    } else {
                        v4 = (1.5 * (3.14159 * (0.0 * v4)))
                    }
                    v31 := GenS2{F0: 2147483647}
                    var v32 bool = (1.5 > 0.1)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case 4294967295:
                var v33 int64 = 1
                _ = v33
            case 2147483647:
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
            }
        default:
            var v34 uint16 = (v5 & v5)
            _ = v34
        }
        _ = fi25
    }
    for fi35 := 0; fi35 < 15; fi35++ {
        delete(v1, 396)
        var v36 uint = 18446744073709551615
        v37 := v1[1]
        _ = fi35
        _ = v36
        _ = v37
    }
    v11 = false
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    v4 -= ((1e10 - float64(float32((v4 / v4)))) / v4)
    var v38 int8 = 100
    for fi39 := 0; fi39 < 17; fi39++ {
        for fi40 := 0; fi40 < 17; fi40++ {
            fmt.Println(fi40, v10)
            v41 := GenS2{F0: 42}
            _ = fi40
            _ = v41
        }
        if v11 {
            v42 := GenS1{F0: -128}
            if len(v16) > 0 {
                v16[0] = fi39
            }
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
            _ = v42
        }
        var v43 uint8 = 1
        _ = fi39
        _ = v43
    }
    v16 = append(v16, -181)
    fmt.Println(v3, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
