package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int {
    v7 := GenS1{F0: -1}
    var v8 uint8 = 21
    _ = v7
    _ = v8
    return 32767
}
func genFn1(p0 uint32, p1 int32) string {
    if ((!((-2.25 + (((3.14159 / (-2.25 * (0.1 / 1.5))) - (0.1 * -2.25)) + 1.5)) >= 0.0)) && false) {
        switch "" {
        case "a":
            p0 -= p0
        case "hello":
            v9 := GenS1{F0: 1}
            _ = v9
        }
    }
    var v10 uint = 4294967295
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            if false {
                v13 := GenS1{F0: 1}
                p0 -= 42
                _ = v13
            } else {
                v14 := []int{1, fi12, -1, genFn0()}
                var v15 uint64 = 2
                _ = v14
                _ = v15
            }
            _ = fi12
        }
        v16 := []int{}
        _ = fi11
        _ = v16
    }
    switch ((("" + ("" + ("" + "\n\t"))) + "foo bar") + "a") {
    case "\n\t":
        var v17 bool = (0.1 >= (1.5 + float64(0.0)))
        _ = v17
    case "a":
        if false {
            switch 849 {
            case -1:
                v18 := []int{4294967295, -933, 9223372036854775807}
                _ = v18
            case 0:
                p1 += (p1 + p1)
            case 2147483647:
                if true {
                    v19 := []int{1, 9223372036854775807, 128, 0}
                    fmt.Println(p1)
                    v10 = 4294967295
                    _ = v19
                } else {
                    var v20 uint8 = 0
                    _ = v20
                }
            }
            p1 -= (p1 - p1)
        } else {
            p0 = p0
            var v21 float64 = 0.0
            p0 = p0
            _ = v21
        }
    case "hello":
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 uint8 = 111
            if (!true) {
                for fi24 := 0; fi24 < 10; fi24++ {
                    if (!false) {
                        switch fi22 {
                        case 1:
                            var v25 int = (fi24 * (fi24 | fi24))
                            _ = v25
                        case 968:
                            fmt.Println(v10, p1, fi22)
                        default:
                            var v26 string = (("hello" + "a") + ("\n\t" + "foo bar"))
                            _ = v26
                        }
                        var v27 uint16 = 675
                        v28 := []int{fi22}
                        _ = v27
                        _ = v28
                    } else {
                        v29 := []int{32767}
                        if (!(!false)) {
                            var v30 int64 = -9223372036854775808
                            if len(v29) >= 2 {
                                v29 = v29[1:len(v29)]
                            }
                            v23 = (^(v23))
                            _ = v30
                        } else {
                            var v31 int8 = 1
                            _ = len(v29)
                            _ = v31
                        }
                        _ = v29
                    }
                    _ = fi24
                }
            }
            _ = fi22
            _ = v23
        }
    default:
        p0 += (p0 ^ p0)
    }
    _ = p0
    _ = p1
    _ = v10
    return "foo bar"
}

func main() {
    var v1 int32 = 0
    var v2 int = 26
    var v3 uint32 = 705
    var v4 uint8 = 2
    var v5 uint64 = 338
    var v6 int8 = 8
    var v32 uint16 = 107
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 bool = true
        v32 = v32
        v35 := GenS1{F0: (v6 << ((684) & 63))}
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := []int{genFn0()}
    for fi37 := 0; fi37 < 8; fi37++ {
        v38 := GenS1{F0: int8(v3)}
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 int8 = -128
        v1 -= (v1 - (v1 >> ((42) & 63)))
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        switch ("foo bar" + ("foo bar" + "foo bar")) {
        case "foo bar":
            v36 = append(v36, int(v6))
        case "":
            v5 = v5
        case "a":
            switch ((("hello" + "a") + "foo bar") + ("\n\t" + "a")) {
            case "":
                if len(v36) >= 2 {
                    v36 = v36[1:len(v36)]
                }
            case "foo bar":
                var v42 uint = uint(v5)
                _ = v42
            }
        }
        v36 = append(v36, (fi41 & (fi41 - fi41)))
        _ = fi41
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 int = (v2 & (-(v2)))
        _ = fi43
        _ = v44
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
