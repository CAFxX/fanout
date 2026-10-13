package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ((("\n\t" + ("a" + ("hello" + (("foo bar" + "a") + "foo bar")))) + "foo bar") + ((("\n\t" + "hello") + "") + "foo bar"))
    v3 := []int{1}
    var v4 int16 = 1000
    var v5 uint64 = 18446744073709551615
    v4 -= v4
    v2 = v2
    _ = len(v3)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v6 int32 = -566
    v7 := GenS1{F0: -755}
    fmt.Println(v4, v4, v4)
    for fi8 := 0; fi8 < 11; fi8++ {
        if false {
            v6 = -2147483648
            var v9 string = "foo bar"
            delete(v1, (fi8 / (((fi8 - fi8)) | 1)))
            _ = v9
        } else {
            v10 := GenS1{F0: 0}
            v11 := v1[-1]
            v4 = (v4 >> ((2) & 63))
            _ = v10
            _ = v11
        }
        var v12 int16 = (v4 >> ((65535) & 63))
        _ = fi8
        _ = v12
    }
    v13 := GenS1{F0: 9223372036854775807}
    switch 74 {
    case 1:
        _ = len(v1)
    case 127:
        v14 := GenS1{F0: -9223372036854775808}
        _ = v14
    case 128:
        var v15 uint16 = 0
        _ = v15
    default:
        v16 := GenS1{F0: 0}
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        v5 = (v5 ^ (v5 - (v5 | (^(v5)))))
        if true {
            switch 0 {
            case 2147483647:
                if false {
                    var v18 int8 = int8(v6)
                    _ = v18
                } else {
                    switch v2 {
                    case "foo bar":
                        for fi19 := 0; fi19 < 18; fi19++ {
                            if false {
                                _ = len(v1)
                            } else {
                                switch fi17 {
                                case 0:
                                    v4 -= v4
                                }
                            }
                            v20 := GenS2{F0: 65535}
                            _ = fi19
                            _ = v20
                        }
                    }
                }
            default:
                v21 := GenS3{F0: fi17, F1: (v4 & v4)}
                _ = v21
            }
            var v22 int = fi17
            _ = v22
        } else {
            if (!(2 == 972)) {
                v6 -= 42
                v6 = v6
            } else {
                var v23 int32 = -1
                v4 = v4
                var v24 float64 = 0.1
                _ = v23
                _ = v24
            }
            fmt.Println(v4, v2, v6)
            if true {
                var v25 int8 = 1
                if len(v3) > 0 {
                    v3[0] = (fi17 + 0)
                }
                _ = v25
            }
        }
        _ = fi17
    }
    v2 = v2
    for fi26 := 0; fi26 < 11; fi26++ {
        v4 = 42
        if ((v2 <= (v2 + (v2 + "\n\t"))) && (-5 > fi26)) {
            if (2 <= 2147483647) {
                if ((v5 >> ((42) & 63)) < v5) {
                    var v27 float64 = float64(0.1)
                    var v28 uint64 = v5
                    fmt.Println(v6, v27)
                    _ = v27
                    _ = v28
                }
                var v29 uint32 = 0
                v6 = v6
                _ = v29
            }
            v30 := GenS4{F0: "", F1: v6}
            delete(v1, fi26)
            _ = v30
        }
        var v31 float64 = -2.25
        _ = fi26
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
