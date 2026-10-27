package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
func genFn0(p0 string, p1 int, p2 int8) uint16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 int64 = 459
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1(p0 int, p1 bool, p2 string) int8 {
    for fi7 := 0; fi7 < 14; fi7++ {
        p2 = p2
        _ = fi7
    }
    v8 := GenS1{F0: 815}
    v9 := GenS1{F0: 18446744073709551615}
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 string = ("\n\t" + "\n\t")
        if (97 == 2147483647) {
            p0 = 9223372036854775807
            if p1 {
                var v12 string = v11
                _ = v12
            }
            v11 = "\n\t"
        } else {
            if (!(!(!p1))) {
                v13 := []int{}
                var v14 int = p0
                if ((!(2147483647 < 2)) || p1) {
                    v11 = ""
                } else {
                    _ = len(v13)
                    for fi15 := 0; fi15 < 11; fi15++ {
                        switch "foo bar" {
                        case "hello":
                            var v16 string = ("" + ((((v11 + ((v11 + ("hello" + p2)) + ((v11 + p2) + (p2 + p2)))) + "\n\t") + (v11 + "foo bar")) + "\n\t"))
                            _ = v16
                        case "":
                            if len(v13) > 0 {
                                v13[0] = fi10
                            }
                        case "foo bar":
                            v13 = append(v13, (^(fi10)))
                        }
                        switch 1 {
                        case -145:
                            var v17 int16 = -32768
                            _ = v17
                        default:
                            p2 = (p2 + v11)
                        }
                        _ = fi15
                    }
                }
                _ = v13
                _ = v14
            } else {
                switch -919 {
                case 32767:
                    v11 = p2
                }
                var v18 uint = 0
                switch fi10 {
                case 898:
                    var v19 string = ("a" + ((((v11 + v11) + "hello") + p2) + v11))
                    _ = v19
                case 32768:
                    var v20 int64 = -152
                    _ = v20
                case 2147483647:
                    v21 := []int{(fi10 * (fi10 - (^(fi10)))), fi10, fi10, p0}
                    _ = v21
                default:
                    if (p1 || (0 < 42)) {
                        var v22 uint64 = 18446744073709551615
                        v23 := GenS1{F0: v18}
                        p1 = (-938 != -2147483648)
                        _ = v22
                        _ = v23
                    } else {
                        v24 := []int{(^(p0)), 9223372036854775807, fi10}
                        _ = v24
                    }
                }
                _ = v18
            }
            var v25 int = fi10
            _ = v25
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 float64 = (1e10 + (3.14159 + 0.0))
    var v2 int64 = 1
    var v3 int = 680
    v2 += v2
    v26 := GenS1{F0: 255}
    var v27 uint8 = 255
    fmt.Println(v2)
    v28 := GenS1{F0: 393}
    switch ("" + ("foo bar" + (((("" + (("foo bar" + "\n\t") + "foo bar")) + ("a" + "foo bar")) + "") + (("" + "foo bar") + "foo bar")))) {
    case "foo bar":
        fmt.Println(v1, v3)
    }
    fmt.Println(v1, v2, v27)
    var v29 string = ("a" + "foo bar")
    var v30 int32 = -2147483648
    var v31 uint = 245
    if false {
        v29 = "a"
    } else {
        v32 := []int{v3, v3, (-(v3))}
        fmt.Println(v31, v2)
        v31 = 204
        _ = v32
    }
    var v33 float64 = v1
    var v34 int16 = 1
    for fi35 := 0; fi35 < 1; fi35++ {
        fmt.Println(v3)
        v36 := []int{fi35, (v3 % ((-646) | 1)), (v3 & 1)}
        _ = fi35
        _ = v36
    }
    v37 := []int{v3, v3}
    switch (^(v3)) {
    case 460:
        if ((v31 == (v31 + v31)) || true) {
            for fi38 := 0; fi38 < 3; fi38++ {
                for fi39 := 0; fi39 < 1; fi39++ {
                    switch 869 {
                    case 4294967295:
                        var v40 int32 = v30
                        _ = v40
                    case 1:
                        v27 = v27
                    }
                    switch v29 {
                    case "foo bar":
                        v41 := GenS2{F0: (2 == 2147483647)}
                        _ = v41
                    }
                    _ = fi39
                }
                _ = fi38
            }
            var v42 string = v29
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v37
}
