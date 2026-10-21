package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint32; F1 string }
func genFn0() uint {
    var v6 int8 = 10
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint64 = 18446744073709551615
        var v9 int32 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    if (42 != 2) {
        v6 = 42
        v6 -= (v6 & (-(v6)))
    } else {
        if false {
            for fi10 := 0; fi10 < 19; fi10++ {
                var v11 int8 = (v6 >> ((42) & 63))
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint32 = 971
                    _ = fi12
                    _ = v13
                }
                _ = fi10
                _ = v11
            }
            var v14 uint32 = 0
            for fi15 := 0; fi15 < 19; fi15++ {
                if true {
                    var v16 uint16 = 0
                    _ = v16
                }
                var v17 int32 = -1
                switch "hello" {
                case "\n\t":
                    var v18 string = ""
                    _ = v18
                case "":
                    v19 := GenS1{F0: 0, F1: 65535, F2: 4294967295}
                    _ = v19
                }
                _ = fi15
                _ = v17
            }
            _ = v14
        }
        v20 := []int{1}
        _ = v20
    }
    v6 = v6
    if true {
        for fi21 := 0; fi21 < 11; fi21++ {
            switch (fi21 | 42) {
            case 9223372036854775807:
                fmt.Println(fi21)
            }
            _ = fi21
        }
        var v22 uint16 = 65535
        _ = v22
    } else {
        if (-1 > 9223372036854775807) {
            v6 -= v6
        } else {
            var v23 int8 = 1
            var v24 uint64 = 42
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: 42}
        var v26 uint16 = 2
        _ = v25
        _ = v26
    }
    _ = v6
    return 2
}

func main() {
    var v1 uint32 = 364
    var v2 uint16 = 654
    v3 := []uint32{(v1 / (((v1 - (^(v1)))) | 1)), (^(v1))}
    var v4 int = 32767
    v5 := make(map[uint32]string)
    if len(v3) > 0 {
        v3[0] = 65535
    }
    fmt.Println(v2, v2, v4)
    _ = len(v5)
    v27 := GenS1{F0: 1, F1: (v2 & 32767), F2: 711}
    v28 := GenS3{F0: -1, F1: 1}
    var v29 int16 = 42
    for fi30 := 0; fi30 < 9; fi30++ {
        for fi31 := 0; fi31 < 8; fi31++ {
            var v32 float32 = 1.5
            v33 := GenS3{F0: 2147483647, F1: 42}
            delete(v5, 1)
            _ = fi31
            _ = v32
            _ = v33
        }
        for fi34 := 0; fi34 < 12; fi34++ {
            switch ("foo bar" + "a") {
            case "\n\t":
                v2 -= (v2 | (^(v2)))
            case "foo bar":
                var v35 int8 = 106
                _ = v35
            case "a":
                if ("hello" > "\n\t") {
                    var v36 bool = true
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v36
                } else {
                    fmt.Println(fi34, v29)
                    v37 := v5[uint32(v2)]
                    v38 := GenS3{F0: 42, F1: 100}
                    _ = v37
                    _ = v38
                }
                v2 = (v2 - v2)
            default:
                v39 := GenS4{F0: (v1 | (v1 % (((v1 >> ((65535) & 63))) | 1))), F1: "a"}
                _ = v39
            }
            _ = fi34
        }
        _ = fi30
    }
    fmt.Println(v2)
    for fi40 := 0; fi40 < 6; fi40++ {
        v29 -= -257
        switch "foo bar" {
        case "a":
            _ = len(v5)
        }
        var v41 string = ("\n\t" + ("hello" + ("a" + ("hello" + (("a" + "foo bar") + "")))))
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
