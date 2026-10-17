package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
func genFn0() int64 {
    v6 := []int{-426, 319, 32768, -51}
    for fi7 := 0; fi7 < 14; fi7++ {
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if (!(366 < 0)) {
                var v8 int32 = -524
                _ = v8
            } else {
                for fi9 := 0; fi9 < 9; fi9++ {
                    var v10 string = ((("foo bar" + ("a" + "foo bar")) + "\n\t") + "a")
                    _ = len(v6)
                    _ = fi9
                    _ = v10
                }
                if true {
                    var v11 uint64 = 18446744073709551615
                    var v12 bool = false
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint8 = 42
                    _ = fi13
                    _ = v14
                }
            }
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 string = ("hello" + "a")
            if true {
                var v17 int = fi15
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        var v18 float64 = 0.1
        _ = fi7
        _ = v18
    }
    if (1e10 <= float64((1.5 * 1.5))) {
        if len(v6) > 0 {
            v6[0] = 127
        }
        v19 := GenS1{F0: 42, F1: 1}
        _ = v19
    }
    switch "a" {
    case "\n\t":
        var v20 int32 = -2147483648
        _ = v20
    case "a":
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 float32 = (1e10 + 1e10)
            for fi23 := 0; fi23 < 15; fi23++ {
                switch fi23 {
                case -952:
                    fmt.Println(v22)
                }
                if len(v6) > 0 {
                    v6[0] = fi23
                }
                fmt.Println(fi21, fi23, fi23)
                _ = fi23
            }
            v22 += float32((float64(v22) * -2.25))
            _ = fi21
            _ = v22
        }
    case "hello":
        var v24 float32 = 1e10
        _ = v24
    }
    _ = v6
    return -583
}

func main() {
    var v1 int64 = -499
    var v2 int = -1
    v3 := []string{}
    var v4 int8 = 1
    v5 := make(map[int]uint32)
    var v25 uint = 0
    fmt.Println(v4, v4, v25)
    switch "hello" {
    case "foo bar":
        v1 += genFn0()
    default:
        var v26 int32 = 0
        _ = v26
    }
    var v27 uint16 = 255
    var v28 int8 = v4
    var v29 bool = (-9223372036854775808 >= v1)
    var v30 uint32 = 2147483647
    for fi31 := 0; fi31 < 1; fi31++ {
        delete(v5, fi31)
        _ = fi31
    }
    var v32 uint = (v25 >> ((v25) & 63))
    if v29 {
        v28 = 76
        if ((-218 != -2147483648) && (uint8(v27) > 127)) {
            var v33 uint16 = (v27 + (v27 % (((v27 % ((42) | 1))) | 1)))
            v34 := GenS1{F0: 134, F1: -634}
            if true {
                v33 += v27
            }
            _ = v33
            _ = v34
        } else {
            v25 = v25
            if len(v3) > 0 {
                v3[0] = ""
            }
            switch v2 {
            case -9223372036854775808:
                v35 := GenS1{F0: -32768, F1: 1}
                _ = v35
            case 42:
                var v36 int = v2
                _ = v36
            case -95:
                if (0 == -2147483648) {
                    fmt.Println(v4, v29)
                } else {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        var v38 bool = v29
                        switch (fi37 >> (((v32 - (^(v32)))) & 63)) {
                        case 247:
                            var v39 uint8 = 2
                            _ = v39
                        case 128:
                            v28 = v28
                        }
                        _ = fi37
                        _ = v38
                    }
                }
            }
        }
        if (!(772 < 673)) {
            var v40 uint64 = 42
            fmt.Println(v1, v27, v2)
            _ = v40
        }
    }
    v5[4294967295] = v30
    var v41 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v41
    _ = v5
}
