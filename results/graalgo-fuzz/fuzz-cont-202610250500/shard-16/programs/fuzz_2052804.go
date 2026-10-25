package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int; F1 int8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 1
    var v3 uint32 = 2147483648
    if true {
        fmt.Println(v3)
    }
    v4 := GenS1{F0: "foo bar", F1: v3}
    var v5 int32 = -703
    fmt.Println(v5, v2)
    for fi6 := 0; fi6 < 3; fi6++ {
        v3 = 4294967295
        _ = fi6
    }
    var v7 uint = 176
    fmt.Println(v2, v7, v7)
    v7 = (v7 - 2)
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{(-(v2))}
        var v10 uint16 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = v2
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(v3, v7, v7)
        _ = fi11
    }
    v12 := GenS1{F0: "", F1: (v3 ^ (v3 & 42))}
    switch (v2 ^ 1) {
    case 32767:
        if ((v5 & 300) >= (^(v5))) {
            if ((0.1 <= 1.5) || false) {
                var v13 uint32 = (v3 & 4294967295)
                v13 = 42
                _ = v13
            } else {
                v14 := GenS2{F0: 0}
                var v15 uint64 = uint64(v3)
                _ = v14
                _ = v15
            }
            var v16 float64 = 1e10
            var v17 uint = (v7 | (v7 * (v7 & 1)))
            _ = v16
            _ = v17
        } else {
            var v18 uint8 = 255
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 uint32 = 0
                v21 := GenS3{F0: 466}
                if false {
                    v22 := GenS1{F0: "hello", F1: v3}
                    _ = v22
                }
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v18
        }
    default:
        v2 = v2
    }
    switch v2 {
    case 9223372036854775807:
        v23 := []int{v2, (v2 | (v2 << (((v7 ^ (v7 * v7))) & 63))), v2, (v2 % ((v2) | 1))}
        _ = v23
    case 504:
        for fi24 := 0; fi24 < 8; fi24++ {
            v25 := GenS3{F0: -32768}
            var v26 string = "\n\t"
            var v27 bool = (!(!(v26 >= (v26 + "\n\t"))))
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
    case 1:
        v28 := GenS4{F0: ("hello" + ("foo bar" + ("\n\t" + ("a" + ""))))}
        _ = v28
    default:
        var v29 string = "a"
        _ = v29
    }
    switch "a" {
    case "\n\t":
        var v30 uint64 = 2
        _ = v30
    case "a":
        switch v2 {
        case 0:
            if true {
                var v31 bool = false
                _ = v31
            } else {
                fmt.Println(v2, v7, v5)
                var v32 string = ""
                v33 := GenS5{F0: 1, F1: 753}
                _ = v32
                _ = v33
            }
        case 2147483647:
            delete(v1, v3)
        case 32768:
            for fi34 := 0; fi34 < 2; fi34++ {
                switch -993 {
                case -715:
                    v3 = v3
                case 42:
                    var v35 int32 = v5
                    _ = v35
                default:
                    if false {
                        var v36 int32 = (v5 - v5)
                        v37 := v1[v3]
                        _ = v36
                        _ = v37
                    }
                }
                var v38 uint16 = 0
                if false {
                    var v39 uint16 = uint16(v3)
                    v2 = (fi34 | v2)
                    _ = v39
                }
                _ = fi34
                _ = v38
            }
        }
    case "foo bar":
        var v40 int16 = -32768
        _ = v40
    }
    switch ((("\n\t" + "a") + "") + (("" + "\n\t") + "")) {
    case "foo bar":
        v3 = v3
    default:
        for fi41 := 0; fi41 < 10; fi41++ {
            v1[v3] = (-(v2))
            v42 := GenS6{F0: (-(v2)), F1: 1}
            _ = fi41
            _ = v42
        }
    }
    var v43 bool = ((v7 - v7) == 255)
    if ((-2.25 != -2.25) && (int8(v2) == 27)) {
        var v44 int = v2
        _ = v44
    } else {
        v3 = v3
        var v45 uint16 = 32767
        fmt.Println(v45)
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v7
}
