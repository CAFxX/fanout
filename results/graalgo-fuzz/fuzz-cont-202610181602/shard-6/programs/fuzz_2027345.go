package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 int32; F1 uint; F2 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int; F1 int16; F2 int }

func main() {
    var v1 float64 = ((float64(1e10) + 1.5) - 0.0)
    var v2 int16 = 0
    v3 := make(map[uint32]int)
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 18; fi6++ {
        v1 -= 1.5
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 bool = false
            _ = fi7
            _ = v8
        }
        v2 = 42
        _ = fi6
    }
    v9 := GenS1{F0: ""}
    for fi10 := 0; fi10 < 4; fi10++ {
        v1 -= v1
        if (-1 <= 1) {
            fmt.Println(fi10, fi10, v2)
            v11 := GenS1{F0: ""}
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int16 = -1
                v2 = (v2 << ((65535) & 63))
                v14 := GenS2{F0: 2}
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        var v15 string = ""
        _ = fi10
        _ = v15
    }
    var v16 uint = 0
    v17 := GenS3{F0: 4294967295, F1: 255}
    v18 := GenS4{F0: -1, F1: (v16 << (((^(v16))) & 63)), F2: 9223372036854775807}
    v2 += v2
    v19 := GenS4{F0: 455, F1: (^(v16)), F2: 42}
    switch 32767 {
    case 32767:
        v2 -= v2
    case -1:
        switch 9223372036854775807 {
        case 32767:
            fmt.Println(v1, v2)
        case 763:
            v16 += (v16 << ((v16) & 63))
        }
    case -2147483648:
        if true {
            if true {
                v20 := GenS1{F0: ""}
                var v21 uint = (v16 | (^(v16)))
                v5[2] = 65535
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS1{F0: ("a" + ("foo bar" + "a"))}
                    _ = fi22
                    _ = v23
                }
                if (0 < 589) {
                    v1 = (((float64(3.14159) - v1) * -2.25) + -2.25)
                    v2 -= (v2 - v2)
                    var v24 int32 = 0
                    _ = v24
                }
            }
            v25 := v5[4294967295]
            if (!(31 <= 2)) {
                var v26 int16 = v2
                switch 32767 {
                case 128:
                    var v27 uint64 = 155
                    _ = v27
                case 2147483647:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        v29 := v3[v25]
                        v25 = 2147483647
                        switch ("foo bar" + "") {
                        case "a":
                            var v30 uint64 = 107
                            _ = v30
                        case "foo bar":
                            var v31 uint64 = 196
                            _ = v31
                        case "":
                            v29 = v29
                        default:
                            fmt.Println(v29)
                        }
                        _ = fi28
                        _ = v29
                    }
                }
                var v32 uint64 = 2
                _ = v26
                _ = v32
            } else {
                v33 := []int{1}
                var v34 int32 = -2147483648
                v2 = 1000
                _ = v33
                _ = v34
            }
            _ = v25
        }
    default:
        v35 := GenS5{F0: 1}
        _ = v35
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        v16 = 821
        v37 := GenS6{F0: (fi36 * fi36), F1: (v2 - v2), F2: (fi36 & 128)}
        _ = fi36
        _ = v37
    }
    if false {
        switch "a" {
        case "":
            var v38 int64 = 483
            _ = v38
        case "foo bar":
            fmt.Println(v2, v1, v16)
        default:
            switch ("a" + "foo bar") {
            case "\n\t":
                v1 = v1
            }
        }
    }
    v2 -= 1
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 bool = false
        switch "foo bar" {
        case "hello":
            v41 := []int{fi39, (fi39 ^ (fi39 + (fi39 / (((fi39 * -1)) | 1)))), fi39}
            _ = v41
        case "":
            v42 := v3[2147483647]
            _ = v42
        case "foo bar":
            v40 = v40
        default:
            var v43 float32 = float32(v1)
            _ = v43
        }
        _ = fi39
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
