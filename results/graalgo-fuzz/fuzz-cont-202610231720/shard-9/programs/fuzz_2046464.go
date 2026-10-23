package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 string }
func genFn0(p0 int32, p1 int16) bool {
    p1 = p1
    if (0 == 42) {
        var v4 float32 = 0.0
        _ = v4
    }
    for fi5 := 0; fi5 < 18; fi5++ {
        switch ("\n\t" + "hello") {
        case "":
            v6 := GenS1{F0: false}
            _ = v6
        case "hello":
            fmt.Println(p1, fi5)
        case "a":
            v7 := GenS1{F0: true}
            _ = v7
        default:
            var v8 bool = (2 == 255)
            _ = v8
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    return (4294967295 > 32768)
}
func genFn1(p0 uint64, p1 int64, p2 int32) uint {
    v9 := GenS1{F0: (!false)}
    for fi10 := 0; fi10 < 12; fi10++ {
        p0 -= (^(p0))
        _ = fi10
    }
    var v11 uint16 = 676
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := []uint32{2147483647}
    v2 := make(map[uint32]string)
    var v3 int8 = -1
    fmt.Println(v3)
    for fi12 := 0; fi12 < 20; fi12++ {
        v1 = append(v1, 4294967295)
        _ = fi12
    }
    v3 = v3
    switch 32767 {
    case 128:
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 14; fi14++ {
                switch (fi14 + (fi13 * -217)) {
                case 386:
                    if len(v1) > 0 {
                        v1[0] = 2
                    }
                case 263:
                    if true {
                        var v15 int8 = 43
                        if (!(18446744073709551615 >= 344)) {
                            fmt.Println(v15)
                            if true {
                                v16 := GenS2{F0: -887}
                                _ = v16
                            } else {
                                var v17 int = (^(fi13))
                                switch "a" {
                                case "foo bar":
                                    if (true && (!(4294967295 > 110))) {
                                        v18 := GenS2{F0: 879}
                                        _ = v18
                                    } else {
                                        switch -1 {
                                        case -187:
                                            fmt.Println(fi13, v15)
                                        case 4294967295:
                                            var v19 float32 = 1.5
                                            _ = v19
                                        case 42:
                                            var v20 int8 = v3
                                            _ = v20
                                        }
                                        var v21 bool = (1 < int32(v17))
                                        _ = v21
                                    }
                                default:
                                    fmt.Println(v15, v17)
                                }
                                _ = v17
                            }
                            v22 := GenS1{F0: (!false)}
                            _ = v22
                        }
                        if false {
                            switch 2147483647 {
                            case 1:
                                var v23 uint8 = 74
                                _ = v23
                            default:
                                fmt.Println(fi13)
                            }
                            v3 = 1
                            var v24 uint8 = 127
                            _ = v24
                        }
                        _ = v15
                    } else {
                        v3 = 0
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                        v3 -= 126
                    }
                case 32768:
                    v3 = (v3 + (-(v3)))
                default:
                    var v25 uint32 = 1
                    _ = v25
                }
                v3 += (v3 / (((^(v3))) | 1))
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                v3 += v3
                _ = fi14
            }
            v26 := GenS3{F0: "hello"}
            _ = fi13
            _ = v26
        }
    case 0:
        switch ((("\n\t" + (("foo bar" + "") + ("a" + ""))) + ((("" + "hello") + "foo bar") + ("foo bar" + "foo bar"))) + (("foo bar" + "") + ("" + ("a" + ("hello" + "hello"))))) {
        case "":
            v3 = (v3 | (^(v3)))
        case "a":
            var v27 string = ((("a" + "hello") + "\n\t") + (("a" + ("hello" + (("hello" + "\n\t") + ("a" + ("\n\t" + ("hello" + "foo bar")))))) + "hello"))
            _ = v27
        case "\n\t":
            if (!((1 != 255) && (1 < -669))) {
                var v28 int8 = v3
                _ = v28
            } else {
                var v29 int32 = 1
                _ = v29
            }
        default:
            fmt.Println(v3)
        }
    }
    var v30 int = 363
    var v31 int16 = 788
    var v32 int32 = 702
    if len(v1) > 0 {
        v1[0] = 42
    }
    v3 += v3
    v33 := GenS2{F0: 9223372036854775807}
    var v34 string = ("a" + "hello")
    v34 = v34
    var v35 float32 = 0.1
    var v36 uint = genFn1(0, 392, (v32 + v32))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
