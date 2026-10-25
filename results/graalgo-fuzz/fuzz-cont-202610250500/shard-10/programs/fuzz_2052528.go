package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint16, p1 uint8) string {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 17; fi5++ {
            p0 = 216
            p0 -= (^(p0))
            var v6 int = fi4
            _ = fi5
            _ = v6
        }
        p1 = (^(p1))
        for fi7 := 0; fi7 < 14; fi7++ {
            p1 = (p1 >> ((4294967295) & 63))
            var v8 int64 = 734
            _ = fi7
            _ = v8
        }
        _ = fi4
    }
    fmt.Println(p1, p1)
    for fi9 := 0; fi9 < 6; fi9++ {
        switch fi9 {
        case 32767:
            v10 := []int{fi9, fi9, fi9, 1}
            _ = v10
        default:
            p0 += p0
        }
        p1 += 0
        fmt.Println(p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    return ""
}

func main() {
    var v1 uint8 = 127
    var v2 bool = false
    v3 := []string{"hello", "a"}
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 float64 = 1.5
        _ = len(v3)
        var v13 uint32 = 1
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 2, F1: 0}
    var v15 uint64 = 322
    switch "a" {
    case "foo bar":
        v2 = (v1 <= 255)
    }
    if (!v2) {
        for fi16 := 0; fi16 < 18; fi16++ {
            switch ("hello" + ((("foo bar" + "a") + ("a" + "\n\t")) + "a")) {
            case "":
                var v17 int32 = 1
                _ = v17
            case "a":
                if false {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 int8 = 86
                        var v20 float32 = 1.5
                        if v2 {
                            if v2 {
                                var v21 uint16 = 32767
                                var v22 uint8 = (^(v1))
                                fmt.Println(v20, fi16)
                                _ = v21
                                _ = v22
                            } else {
                                v3 = append(v3, "hello")
                                fmt.Println(v1, v19, fi16)
                                var v23 int8 = (v19 ^ 3)
                                _ = v23
                            }
                            var v24 int16 = int16(v1)
                            var v25 uint32 = 690
                            v19 = v19
                            _ = len(v3)
                            v26 := GenS1{F0: 353, F1: (v19 | v19)}
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            v19 -= v19
                        }
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    v1 += (v1 * (v1 >> ((42) & 63)))
                } else {
                    var v27 int8 = 72
                    var v28 int32 = -1
                    _ = v27
                    _ = v28
                }
            default:
                var v29 int64 = -9223372036854775808
                _ = v29
            }
            var v30 int8 = -20
            _ = fi16
            _ = v30
        }
    }
    v2 = v2
    v3 = append(v3, "foo bar")
    if ((!(!true)) && (v1 == 108)) {
        var v31 int32 = 2147483647
        v2 = v2
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
