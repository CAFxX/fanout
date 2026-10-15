package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int32; F2 string }
func genFn0(p0 int32) uint16 {
    var v6 int64 = 1
    for fi7 := 0; fi7 < 14; fi7++ {
        v6 += v6
        v8 := []int{(-(fi7)), fi7, 42}
        v6 = (v6 >> ((65535) & 63))
        _ = fi7
        _ = v8
    }
    var v9 float64 = (float64((1.5 * 1.5)) + 0.1)
    fmt.Println(p0, p0)
    v10 := GenS1{F0: 0, F1: -527}
    _ = p0
    _ = v10
    _ = v6
    _ = v9
    return 32767
}
func genFn1(p0 uint16, p1 string) uint16 {
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint = 36
        v12 += 364
        if (!true) {
            if (true && (true || true)) {
                var v13 float64 = 0.0
                fmt.Println(v13, v12)
                _ = v13
            }
            v14 := []int{fi11}
            _ = v14
        } else {
            if true {
                p1 = ((p1 + ("" + p1)) + "")
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS2{F0: uint32(p0), F1: 934}
                    p1 = ""
                    v17 := []int{}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        }
        _ = fi11
        _ = v12
    }
    v18 := []int{}
    p0 = 440
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v18
    return 864
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = -9223372036854775808
    var v3 int64 = -1
    var v4 int16 = 42
    v5 := make(map[string]int)
    var v19 string = ("foo bar" + "")
    if (400 > 730) {
        if ((1.5 - 3.14159) != 1e10) {
            v20 := GenS1{F0: 126, F1: 9223372036854775807}
            _ = v20
        } else {
            if true {
                var v21 int8 = -91
                _ = v21
            } else {
                var v22 uint8 = 2
                v23 := GenS1{F0: -1, F1: (v2 << ((205) & 63))}
                _ = v22
                _ = v23
            }
            if true {
                if false {
                    if (2147483648 < 2147483647) {
                        v4 += (-(v4))
                        fmt.Println(v3, v3)
                    }
                    v24 := GenS3{F0: int(v3), F1: (v3 << ((42) & 63))}
                    _ = v24
                } else {
                    v5["a"] = (^(v2))
                    v2 += v2
                }
                v1["foo bar"] = 42
                v2 -= 32767
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v4 = (v4 >> ((255) & 63))
                    _ = fi25
                }
            }
            switch v19 {
            case "foo bar":
                var v26 int = -2147483648
                _ = v26
            }
        }
        v27 := GenS4{F0: 0}
        _ = v27
    } else {
        v19 = "foo bar"
    }
    for fi28 := 0; fi28 < 12; fi28++ {
        var v29 int32 = -2147483648
        var v30 int16 = (^(v4))
        v19 = v19
        _ = fi28
        _ = v29
        _ = v30
    }
    v4 = (v4 ^ (v4 << ((42) & 63)))
    var v31 int64 = v3
    v32 := GenS1{F0: 1, F1: v2}
    switch -568 {
    case 127:
        var v33 int64 = (v3 % ((1) | 1))
        _ = v33
    case 512:
        v34 := GenS1{F0: -322, F1: 32767}
        _ = v34
    case 955:
        var v35 bool = (!true)
        _ = v35
    }
    switch (("\n\t" + ((((v19 + ((((v19 + (v19 + "")) + (((v19 + v19) + "\n\t") + v19)) + v19) + "hello")) + "\n\t") + (((v19 + ("\n\t" + "foo bar")) + "\n\t") + v19)) + (v19 + v19))) + (v19 + ("foo bar" + (((v19 + v19) + v19) + v19)))) {
    case "a":
        v36 := GenS5{F0: v2, F1: -471, F2: v19}
        _ = v36
    case "":
        v37 := []int{v2, (v2 ^ (v2 % (((v2 >> ((65535) & 63))) | 1))), 488, v2}
        _ = v37
    case "foo bar":
        fmt.Println(v19, v19)
    default:
        for fi38 := 0; fi38 < 10; fi38++ {
            delete(v5, ((v19 + ("foo bar" + v19)) + (v19 + v19)))
            v1[(v19 + (v19 + "foo bar"))] = 4294967295
            v31 += v3
            _ = fi38
        }
    }
    v39 := []int{v2, v2, 698, v2}
    var v40 uint8 = 15
    v41 := GenS4{F0: 2147483647}
    switch 2147483647 {
    case -183:
        var v42 uint32 = 467
        _ = v42
    case 128:
        v43 := GenS3{F0: v2, F1: -667}
        _ = v43
    case 189:
        var v44 uint8 = (v40 - (v40 / (((v40 & (^(v40)))) | 1)))
        _ = v44
    }
    v40 += (v40 + 128)
    switch (v19 + v19) {
    case "hello":
        for fi45 := 0; fi45 < 2; fi45++ {
            for fi46 := 0; fi46 < 18; fi46++ {
                _ = len(v1)
                var v47 uint8 = v40
                _ = fi46
                _ = v47
            }
            for fi48 := 0; fi48 < 18; fi48++ {
                if false {
                    var v49 string = (v19 + (v19 + ((v19 + v19) + "hello")))
                    var v50 string = v49
                    _ = len(v39)
                    _ = v49
                    _ = v50
                }
                _ = fi48
            }
            _ = fi45
        }
    case "\n\t":
        v31 = 0
    default:
        if false {
            var v51 uint32 = 507
            fmt.Println(v19, v51, v19)
            if (2 > uint8(v51)) {
                fmt.Println(v31)
                v3 -= v3
                if len(v39) >= 2 {
                    v39 = v39[1:len(v39)]
                }
            }
            _ = v51
        }
    }
    v5[("foo bar" + "hello")] = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
