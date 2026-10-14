package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 int; F2 int64 }
func genFn0(p0 bool) int32 {
    v5 := GenS1{F0: 32767, F1: 203}
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    if ((true || (!("hello" > "foo bar"))) && (308 > -75)) {
        var v8 int = 2147483647
        _ = v8
    }
    v9 := GenS1{F0: -1, F1: 1}
    switch ("hello" + "\n\t") {
    case "a":
        p0 = (p0 && false)
    case "\n\t":
        if p0 {
            var v10 int64 = 1
            _ = v10
        }
    case "hello":
        for fi11 := 0; fi11 < 6; fi11++ {
            switch "hello" {
            case "a":
                for fi12 := 0; fi12 < 17; fi12++ {
                    v13 := []int{(-(fi12))}
                    if (4294967295 >= 4294967295) {
                        var v14 int = (fi12 * fi11)
                        _ = v14
                    } else {
                        var v15 uint8 = 0
                        v15 = (v15 & 255)
                        v15 = (v15 % ((v15) | 1))
                        _ = v15
                    }
                    v16 := GenS2{F0: 27, F1: 128}
                    _ = fi12
                    _ = v13
                    _ = v16
                }
            }
            v17 := GenS3{F0: 897}
            v18 := []int{(fi11 + fi11), fi11, 128}
            _ = fi11
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = v5
    _ = v9
    return -1
}

func main() {
    var v1 int32 = 2147483647
    var v2 string = "foo bar"
    var v3 int8 = -1
    v4 := make(map[string]string)
    var v19 int8 = (v3 + (v3 % ((v3) | 1)))
    if false {
        var v20 uint64 = 18446744073709551615
        if false {
            v19 += v19
            v21 := []int{}
            var v22 uint8 = 255
            _ = v21
            _ = v22
        } else {
            switch 128 {
            case 32767:
                fmt.Println(v1, v2, v2)
            case -291:
                var v23 bool = false
                _ = v23
            }
            var v24 uint32 = 869
            _ = v24
        }
        var v25 uint = 18446744073709551615
        _ = v20
        _ = v25
    } else {
        if (65535 >= 1) {
            v2 = (((v2 + v2) + v2) + v2)
            if false {
                switch 128 {
                case 127:
                    var v26 bool = true
                    _ = v26
                case -2147483648:
                    var v27 uint = 18446744073709551615
                    _ = v27
                default:
                    var v28 bool = true
                    _ = v28
                }
            }
        } else {
            if (9223372036854775807 < 9223372036854775807) {
                var v29 uint16 = 988
                fmt.Println(v29, v2, v2)
                _ = v29
            }
            v3 -= (^(v19))
        }
        v30 := []int{4294967295, 0, -503, 32768}
        _ = v30
    }
    switch v2 {
    case "foo bar":
        switch v2 {
        case "a":
            v31 := []int{32768, 335, 4294967295, 9223372036854775807}
            _ = v31
        default:
            var v32 float32 = 1.5
            _ = v32
        }
    default:
        var v33 uint64 = 128
        _ = v33
    }
    if true {
        var v34 uint64 = 18446744073709551615
        if (2 < 241) {
            for fi35 := 0; fi35 < 7; fi35++ {
                var v36 uint = 4294967295
                var v37 int8 = (v19 * -1)
                _ = fi35
                _ = v36
                _ = v37
            }
        }
        var v38 int = 0
        _ = v34
        _ = v38
    }
    v3 = v19
    v2 = (v2 + ((v2 + ("foo bar" + ((((v2 + v2) + ("foo bar" + v2)) + (v2 + "a")) + ("foo bar" + v2)))) + ""))
    var v39 float32 = 1e10
    v40 := GenS4{F0: 58, F1: -1, F2: 9223372036854775807}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
}
