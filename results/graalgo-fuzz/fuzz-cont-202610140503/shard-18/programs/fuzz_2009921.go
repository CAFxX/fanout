package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        switch ("\n\t" + (("foo bar" + ("foo bar" + "a")) + "")) {
        case "hello":
            var v7 int = (fi6 | fi6)
            _ = v7
        default:
            var v8 float64 = 1e10
            _ = v8
        }
        _ = fi6
    }
    var v9 int32 = 182
    v9 = v9
    fmt.Println(v9)
    fmt.Println(v9)
    _ = v9
    return 1
}
func genFn1() uint8 {
    v10 := GenS1{F0: 0}
    if false {
        var v11 uint16 = 32767
        _ = v11
    } else {
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 int64 = -1
            switch "" {
            case "":
                for fi14 := 0; fi14 < 18; fi14++ {
                    switch ("\n\t" + "hello") {
                    case "a":
                        if (!(255 >= 32767)) {
                            var v15 string = ("foo bar" + "")
                            _ = v15
                        }
                    case "hello":
                        var v16 string = ""
                        _ = v16
                    case "":
                        var v17 float64 = -2.25
                        _ = v17
                    default:
                        v13 += 1
                    }
                    var v18 uint = 627
                    fmt.Println(fi14, v13, fi14)
                    _ = fi14
                    _ = v18
                }
            }
            for fi19 := 0; fi19 < 10; fi19++ {
                v13 = 0
                v20 := GenS2{F0: (false || false)}
                _ = fi19
                _ = v20
            }
            _ = fi12
            _ = v13
        }
    }
    switch (("foo bar" + "hello") + "") {
    case "foo bar":
        var v21 float32 = 1e10
        _ = v21
    default:
        var v22 int32 = -546
        _ = v22
    }
    _ = v10
    return 128
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 981
    var v3 uint64 = 18446744073709551615
    var v4 int64 = 638
    var v5 uint8 = 127
    var v23 float32 = 0.0
    v24 := GenS2{F0: (0 == 296)}
    v4 -= v4
    if ((false && true) || (!true)) {
        if false {
            v25 := GenS2{F0: (16 == v5)}
            _ = v25
        }
        if true {
            for fi26 := 0; fi26 < 19; fi26++ {
                v27 := GenS3{F0: uint64(v5)}
                _ = fi26
                _ = v27
            }
        } else {
            var v28 int64 = v4
            v1[446] = -1
            v29 := GenS1{F0: -602}
            _ = v28
            _ = v29
        }
        for fi30 := 0; fi30 < 2; fi30++ {
            if true {
                for fi31 := 0; fi31 < 16; fi31++ {
                    var v32 int32 = 0
                    switch 4294967295 {
                    case -1:
                        var v33 int32 = -433
                        if true {
                            fmt.Println(fi30)
                            v34 := GenS2{F0: false}
                            var v35 uint = 65535
                            _ = v34
                            _ = v35
                        }
                        var v36 string = ""
                        _ = v33
                        _ = v36
                    case -761:
                        v37 := GenS2{F0: false}
                        _ = v37
                    default:
                        var v38 int16 = 32767
                        _ = v38
                    }
                    fmt.Println(v23, v5, v5)
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi30
        }
    }
    var v39 uint64 = v3
    v40 := GenS4{F0: 339}
    switch -9223372036854775808 {
    case -2147483648:
        v4 -= v4
    default:
        v4 = (-(v4))
    }
    v41 := GenS5{F0: 9223372036854775807, F1: (true && (!true))}
    var v42 bool = true
    v42 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
