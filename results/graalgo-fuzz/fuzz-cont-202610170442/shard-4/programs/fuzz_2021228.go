package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
func genFn0(p0 int64, p1 int8, p2 uint16) int32 {
    var v4 uint32 = 65535
    v5 := []int{-717, 0, 552}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -404
}

func main() {
    var v1 float32 = 1.5
    var v2 int8 = 42
    var v3 int64 = -62
    var v6 uint64 = 543
    for fi7 := 0; fi7 < 14; fi7++ {
        v8 := []int{(fi7 & (^(fi7))), fi7}
        var v9 float64 = 0.1
        fmt.Println(v2, v1, v6)
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        if ("foo bar" == "") {
            for fi11 := 0; fi11 < 1; fi11++ {
                switch (fi10 + (fi10 | fi10)) {
                case -64:
                    var v12 float64 = ((-2.25 + 0.1) * 1.5)
                    _ = v12
                case 1:
                    switch -810 {
                    case -825:
                        v13 := []int{}
                        _ = v13
                    case -439:
                        fmt.Println(v6, fi10, fi11)
                    case 127:
                        fmt.Println(v6)
                    }
                    var v14 int32 = 202
                    _ = v14
                case 42:
                    var v15 int32 = 1
                    _ = v15
                default:
                    v16 := GenS1{F0: (fi11 & (fi11 % ((fi11) | 1)))}
                    _ = v16
                }
                fmt.Println(v2, fi11)
                _ = fi11
            }
            v17 := []int{-1, fi10, (fi10 | (fi10 + fi10))}
            _ = v17
        }
        _ = fi10
    }
    v6 += v6
    v6 = v6
    if (!false) {
        fmt.Println(v2, v2, v2)
        v18 := []int{}
        _ = v18
    }
    if (4294967295 != 479) {
        v2 = (v2 - v2)
    } else {
        v19 := []int{127, -1, 0, -9223372036854775808}
        v20 := GenS1{F0: 106}
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 6; fi22++ {
            if false {
                switch (("" + "") + "hello") {
                case "hello":
                    var v23 float32 = (v1 - (v1 / v1))
                    _ = v23
                case "\n\t":
                    v24 := []int{fi22, fi21, (fi22 | fi21)}
                    _ = v24
                }
                if (!((false && true) && (((18446744073709551615 <= (v6 * v6)) || true) && (18446744073709551615 != 255)))) {
                    fmt.Println(v3, v3)
                }
            }
            _ = fi22
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            v3 -= 42
            v26 := GenS2{F0: (fi25 << ((1) & 63))}
            _ = fi25
            _ = v26
        }
        var v27 int64 = (v3 >> ((42) & 63))
        _ = fi21
        _ = v27
    }
    switch -803 {
    case 1:
        v28 := GenS2{F0: 2147483647}
        _ = v28
    case 4294967295:
        v3 = v3
    case 261:
        v29 := []int{121, 2147483647}
        _ = v29
    }
    for fi30 := 0; fi30 < 16; fi30++ {
        fmt.Println(v2)
        _ = fi30
    }
    v31 := []int{42, 4294967295, 42}
    v1 = 3.14159
    v32 := GenS2{F0: 9223372036854775807}
    v2 += v2
    var v33 int = -1
    if (!true) {
        for fi34 := 0; fi34 < 3; fi34++ {
            if (true || true) {
                if (!(255 < 225)) {
                    for fi35 := 0; fi35 < 18; fi35++ {
                        var v36 bool = false
                        _ = fi35
                        _ = v36
                    }
                }
                v6 = (v6 % (((v6 | (v6 - (v6 << ((624) & 63))))) | 1))
            } else {
                var v37 string = "\n\t"
                v33 = (-(fi34))
                _ = v37
            }
            _ = fi34
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            for fi39 := 0; fi39 < 2; fi39++ {
                fmt.Println(v6, v2)
                var v40 uint = 18446744073709551615
                _ = fi39
                _ = v40
            }
            var v41 uint16 = 0
            if true {
                fmt.Println(v2, v1, v41)
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
            } else {
                var v42 uint16 = v41
                var v43 bool = ("a" == (((("foo bar" + "") + "a") + "foo bar") + (((((("foo bar" + ("" + "")) + ("foo bar" + "foo bar")) + "") + (("" + "foo bar") + "a")) + ("" + ("hello" + "a"))) + "\n\t")))
                _ = v42
                _ = v43
            }
            _ = fi38
            _ = v41
        }
    }
    var v44 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v44
    _ = v6
}
