package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint; F1 uint16; F2 uint16 }
func genFn0(p0 string, p1 int16) bool {
    fmt.Println(p0)
    p1 -= p1
    _ = p0
    _ = p1
    return false
}

func main() {
    var v1 int8 = 4
    v2 := make(map[string]bool)
    var v3 int64 = 42
    var v4 int16 = -32768
    v5 := make(map[int]uint32)
    fmt.Println(v1, v3)
    v6 := GenS1{F0: true}
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 uint32 = 2
        _ = fi7
        _ = v8
    }
    v4 = v4
    switch 127 {
    case -9223372036854775808:
        v4 += 42
    case 221:
        _ = len(v2)
    case 0:
        v4 -= v4
    default:
        if true {
            v9 := v2["hello"]
            _ = v9
        }
    }
    if ((true || false) && true) {
        v10 := GenS1{F0: (!false)}
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                switch ("\n\t" + "a") {
                case "foo bar":
                    v13 := []int{fi12, 1, (fi12 & fi11), fi11}
                    _ = v13
                case "a":
                    var v14 uint64 = 661
                    _ = v14
                case "\n\t":
                    fmt.Println(v4, v1, fi11)
                }
                var v15 float64 = 0.0
                _ = fi12
                _ = v15
            }
            if (false && ((!(!(!(fi11 > fi11)))) && (!(("hello" + (((("hello" + "foo bar") + ("" + ("\n\t" + ("" + ("foo bar" + "foo bar"))))) + "foo bar") + "\n\t")) <= "foo bar")))) {
                var v16 int64 = (v3 + (v3 / ((0) | 1)))
                fmt.Println(v1)
                _ = v16
            } else {
                _ = len(v5)
                fmt.Println(v1)
                v4 = v4
            }
            _ = fi11
        }
        v4 = 32767
        _ = v10
    }
    if true {
        var v17 int = 2147483647
        switch 128 {
        case -358:
            for fi18 := 0; fi18 < 2; fi18++ {
                v19 := []int{(v17 << ((97) & 63))}
                if ((true || false) || (false && (true && false))) {
                    v17 = v17
                    var v20 uint8 = 42
                    _ = v20
                }
                v21 := GenS1{F0: true}
                _ = fi18
                _ = v19
                _ = v21
            }
        case 1:
            var v22 int8 = (v1 % ((v1) | 1))
            _ = v22
        }
        v23 := GenS1{F0: ((false || true) || (v4 >= v4))}
        _ = v17
        _ = v23
    }
    v1 = int8(v3)
    v24 := []int{}
    v25 := GenS2{F0: 0, F1: "hello"}
    v3 -= v3
    var v26 float32 = (1e10 / (-2.25 / 0.1))
    for fi27 := 0; fi27 < 13; fi27++ {
        fmt.Println(fi27)
        v3 -= 0
        _ = len(v2)
        _ = fi27
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            fmt.Println(v4, v4)
            v26 -= v26
            _ = fi29
        }
        var v30 uint64 = 514
        if (0.0 != 1e10) {
            v31 := GenS1{F0: ("foo bar" >= "\n\t")}
            switch ("hello" + ("hello" + "")) {
            case "foo bar":
                v32 := GenS1{F0: (!(827 == 637))}
                _ = v32
            }
            _ = v31
        } else {
            var v33 int = (fi28 >> ((0) & 63))
            _ = v33
        }
        _ = fi28
        _ = v30
    }
    v34 := GenS3{F0: 255, F1: 911, F2: 542}
    v26 = v26
    var v35 float32 = ((0.0 + ((0.0 / float32(-2.25)) / 3.14159)) - 1e10)
    v5[int(v4)] = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
