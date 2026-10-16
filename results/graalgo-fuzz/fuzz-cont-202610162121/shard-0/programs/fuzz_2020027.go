package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int }
func genFn0(p0 int64, p1 uint8) uint {
    if (!true) {
        v4 := GenS1{F0: true}
        var v5 string = ""
        if false {
            p1 += 1
            v6 := []int{}
            var v7 uint = 376
            _ = v6
            _ = v7
        }
        _ = v4
        _ = v5
    } else {
        var v8 float32 = 1e10
        var v9 uint16 = 255
        var v10 string = ("a" + ("hello" + "foo bar"))
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(p1, p1)
    var v11 int32 = -2147483648
    switch "foo bar" {
    case "a":
        v12 := GenS1{F0: (true || false)}
        _ = v12
    case "hello":
        v13 := GenS1{F0: (42 <= -1)}
        _ = v13
    case "":
        v14 := GenS2{F0: "\n\t", F1: -9223372036854775808}
        _ = v14
    default:
        if false {
            var v15 int64 = (p0 / ((9223372036854775807) | 1))
            for fi16 := 0; fi16 < 3; fi16++ {
                for fi17 := 0; fi17 < 20; fi17++ {
                    var v18 float32 = 1.5
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return 255
}

func main() {
    var v1 uint64 = 984
    v2 := make(map[string]uint32)
    var v3 int = 335
    v19 := GenS3{F0: 0}
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := []int{}
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS2{F0: ("foo bar" + ("a" + "\n\t")), F1: fi20}
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    switch (-(v3)) {
    case 1:
        v2[("" + "hello")] = 4294967295
    case 0:
        if true {
            v24 := GenS2{F0: "", F1: v3}
            v25 := GenS2{F0: ("" + "foo bar"), F1: -35}
            var v26 int32 = 1
            _ = v24
            _ = v25
            _ = v26
        } else {
            switch "hello" {
            case "\n\t":
                for fi27 := 0; fi27 < 2; fi27++ {
                    if false {
                        if ((!false) && ((-2.25 / 1.5) == -2.25)) {
                            v28 := GenS2{F0: "a", F1: v3}
                            var v29 float64 = (-2.25 / ((1.5 * -2.25) / 1.0))
                            _ = v28
                            _ = v29
                        }
                    } else {
                        switch (v3 | fi27) {
                        case 128:
                            var v30 string = ((("" + ("foo bar" + ("" + ("a" + "")))) + "a") + "")
                            _ = v30
                        case 9223372036854775807:
                            v3 += (v3 + v3)
                        default:
                            v31 := GenS2{F0: ("" + ("\n\t" + "hello")), F1: (fi27 >> ((42) & 63))}
                            _ = v31
                        }
                        v1 = v1
                    }
                    _ = fi27
                }
            case "foo bar":
                v32 := []int{v3, 559, 1, v3}
                _ = v32
            default:
                v1 += 991
            }
            v1 = (v1 & (v1 - (v1 * v1)))
        }
    case 4294967295:
        var v33 uint16 = 241
        _ = v33
    }
    fmt.Println(v3, v1)
    fmt.Println(v3, v1)
    var v34 uint = genFn0(42, 2)
    if (2147483647 <= 1) {
        var v35 uint = (v34 % ((v34) | 1))
        _ = v35
    }
    var v36 uint32 = 1
    var v37 int8 = -35
    switch "hello" {
    case "\n\t":
        var v38 int64 = -1
        _ = v38
    case "":
        if true {
            v39 := GenS4{F0: int(v1)}
            _ = v39
        } else {
            var v40 string = "\n\t"
            _ = v40
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v36
    _ = v37
}
