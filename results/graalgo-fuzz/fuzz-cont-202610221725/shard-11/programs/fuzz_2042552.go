package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int64; F1 int16; F2 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint32; F1 uint16 }
func genFn0(p0 int, p1 uint16) int64 {
    var v6 string = ((("\n\t" + "") + "\n\t") + "\n\t")
    v7 := []int{p0, (p0 / ((1) | 1)), p0, (p0 & p0)}
    v8 := GenS1{F0: 232}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -719
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 100
    v3 := make(map[string]string)
    var v4 int8 = (v2 << ((905) & 63))
    var v5 uint = 18446744073709551615
    fmt.Println(v5, v2, v5)
    fmt.Println(v2, v2)
    v9 := GenS2{F0: (v5 | v5), F1: 127}
    switch ("a" + "") {
    case "hello":
        fmt.Println(v5, v5)
    case "foo bar":
        v10 := []int{-349, 2147483647, 128, 32767}
        _ = v10
    case "a":
        var v11 float32 = 1e10
        _ = v11
    }
    if (!(0.1 != 3.14159)) {
        if ((255 > 863) || true) {
            v12 := GenS2{F0: 4294967295, F1: 74}
            v13 := []int{-148, 32767}
            _ = v12
            _ = v13
        } else {
            delete(v1, "")
        }
    } else {
        if ((v5 << (((v5 | v5)) & 63)) == (v5 << ((v5) & 63))) {
            switch -1 {
            case 4294967295:
                v14 := GenS3{F0: 1, F1: -32768, F2: 0}
                _ = v14
            }
            if (true || (954 == 2)) {
                var v15 int32 = 42
                _ = v15
            } else {
                v16 := GenS1{F0: 2147483648}
                _ = v16
            }
        } else {
            var v17 uint = v5
            _ = v17
        }
        if (!(!(643 < v5))) {
            switch 9223372036854775807 {
            case 127:
                var v18 bool = false
                _ = v18
            case 1:
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 bool = (!(29 < uint64(v5)))
                    _ = fi19
                    _ = v20
                }
            }
            if false {
                v21 := GenS2{F0: 1, F1: 128}
                v22 := GenS4{F0: 65535, F1: 37, F2: -9223372036854775808}
                v23 := GenS2{F0: (v5 & v5), F1: 129}
                _ = v21
                _ = v22
                _ = v23
            }
            v24 := GenS3{F0: 42, F1: -1, F2: 42}
            _ = v24
        }
    }
    if ((((("a" + "hello") + ("hello" + "a")) + "") + "a") >= "") {
        v25 := v3["foo bar"]
        if (v2 <= v4) {
            switch 32768 {
            case 4294967295:
                v26 := GenS2{F0: 42, F1: 127}
                _ = v26
            case 1:
                v25 = "hello"
            default:
                for fi27 := 0; fi27 < 10; fi27++ {
                    if ((v4 ^ (v4 << ((v5) & 63))) <= v4) {
                        switch v25 {
                        case "":
                            v4 = v4
                        case "hello":
                            fmt.Println(v4)
                        }
                    }
                    if false {
                        v28 := GenS5{F0: v2, F1: false}
                        var v29 int16 = 1000
                        if false {
                            for fi30 := 0; fi30 < 17; fi30++ {
                                v29 -= v29
                                var v31 uint64 = 2
                                v32 := GenS6{F0: 1, F1: 2}
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                        _ = v28
                        _ = v29
                    } else {
                        fmt.Println(v2, v4)
                        var v33 int8 = v4
                        _ = v33
                    }
                    if true {
                        v4 -= 105
                    }
                    _ = fi27
                }
            }
            var v34 string = ((((("a" + "hello") + "hello") + "") + v25) + (("foo bar" + "hello") + v25))
            v25 = (v25 + v34)
            _ = v34
        } else {
            v2 = (v4 & v2)
            v4 = v4
        }
        delete(v1, "foo bar")
        _ = v25
    } else {
        v35 := GenS6{F0: 65535, F1: 255}
        _ = v35
    }
    v4 = v4
    var v36 int16 = 339
    v37 := GenS1{F0: 2}
    v4 = v4
    switch "\n\t" {
    case "":
        var v38 int64 = -879
        _ = v38
    case "hello":
        for fi39 := 0; fi39 < 5; fi39++ {
            v40 := []int{fi39, -174, (fi39 - 9223372036854775807), fi39}
            _ = fi39
            _ = v40
        }
    default:
        for fi41 := 0; fi41 < 8; fi41++ {
            for fi42 := 0; fi42 < 9; fi42++ {
                v43 := []int{fi41}
                _ = fi42
                _ = v43
            }
            _ = fi41
        }
    }
    v36 = 0
    var v44 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v44
    _ = v5
    _ = v9
}
