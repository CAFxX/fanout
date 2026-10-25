package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0(p0 bool, p1 uint8) string {
    for fi4 := 0; fi4 < 1; fi4++ {
        switch (^(fi4)) {
        case 0:
            var v5 float64 = -2.25
            _ = v5
        case 496:
            p0 = false
        }
        if p0 {
            p0 = p0
            var v6 int = fi4
            _ = v6
        }
        _ = fi4
    }
    v7 := GenS1{F0: ((-2.25 * 3.14159) < 3.14159), F1: 0, F2: 42}
    v8 := []int{}
    var v9 bool = p0
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + "foo bar")
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v10 float32 = 3.14159
    var v11 string = "\n\t"
    fmt.Println(v10)
    v12 := GenS2{F0: 65535, F1: 255}
    var v13 uint32 = 65535
    v14 := GenS1{F0: ((((false || (true || true)) || true) || (!(v11 != ((v11 + v11) + v11)))) && (99 >= 42)), F1: 9223372036854775807, F2: 675}
    if true {
        switch 128 {
        case 42:
            if true {
                v11 = (v11 + ((v11 + (v11 + "hello")) + "\n\t"))
                var v15 uint16 = 32767
                _ = v15
            } else {
                v10 = float32((-2.25 / 1e10))
                var v16 int = 342
                _ = v16
            }
        case -2147483648:
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS1{F0: (2 > 42), F1: 127, F2: 371}
                if (!((v13 ^ 65535) <= (v13 * 905))) {
                    v10 -= 1.5
                    for fi19 := 0; fi19 < 17; fi19++ {
                        v20 := GenS3{F0: -1, F1: 0, F2: 0}
                        _ = fi19
                        _ = v20
                    }
                    var v21 int32 = -1
                    _ = v21
                } else {
                    var v22 bool = true
                    _ = v22
                }
                v11 = v11
                _ = fi17
                _ = v18
            }
        case 0:
            fmt.Println(v10, v13)
        default:
            for fi23 := 0; fi23 < 10; fi23++ {
                delete(v3, ((v11 + v11) + ((v11 + v11) + "")))
                var v24 uint8 = 255
                _ = fi23
                _ = v24
            }
        }
        v25 := GenS3{F0: 32767, F1: 1, F2: 18446744073709551615}
        _ = v25
    }
    v26 := []int{-9223372036854775808}
    switch (v11 + v11) {
    case "hello":
        switch ("a" + v11) {
        case "hello":
            v13 = (v13 & v13)
        case "a":
            fmt.Println(v11, v10, v11)
        case "foo bar":
            if false {
                for fi27 := 0; fi27 < 14; fi27++ {
                    if true {
                        for fi28 := 0; fi28 < 4; fi28++ {
                            var v29 uint16 = 221
                            v13 = 2147483648
                            switch "foo bar" {
                            case "hello":
                                var v30 int32 = -954
                                _ = v30
                            default:
                                fmt.Println(v29, fi27, fi27)
                                v13 += 65535
                                v13 += v13
                            }
                            _ = fi28
                            _ = v29
                        }
                        fmt.Println(fi27, v10, v10)
                        fmt.Println(v11, v13)
                    }
                    var v31 float32 = 3.14159
                    v32 := GenS3{F0: 714, F1: 264, F2: 42}
                    _ = fi27
                    _ = v31
                    _ = v32
                }
                v10 += v10
            }
        }
    case "":
        if true {
            var v33 int32 = 42
            switch "\n\t" {
            case "":
                fmt.Println(v10, v13, v13)
            }
            _ = v33
        }
    case "foo bar":
        var v34 uint16 = 0
        _ = v34
    default:
        v13 = v13
    }
    var v35 string = (v11 + ("foo bar" + "\n\t"))
    var v36 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v36
}
