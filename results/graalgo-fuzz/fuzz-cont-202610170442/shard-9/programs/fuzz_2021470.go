package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint16, p2 int) uint {
    switch "" {
    case "\n\t":
        if ((!false) || (4294967295 == 594)) {
            v7 := GenS1{F0: 509}
            var v8 uint8 = p0
            _ = v7
            _ = v8
        } else {
            var v9 uint8 = p0
            fmt.Println(p0, p0, p0)
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := []int{(p2 >> ((883) & 63)), p2, (p2 | 0), p2}
                p2 = (fi10 ^ fi10)
                var v12 int64 = -1
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = v9
        }
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 int = -631
        var v15 uint64 = 1
        var v16 uint16 = 889
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 965
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483648
    v5 := []string{(("\n\t" + "a") + ""), (((("\n\t" + ("hello" + ("\n\t" + (((("hello" + (("" + "a") + "foo bar")) + "a") + ("\n\t" + "\n\t")) + ("a" + "\n\t"))))) + "foo bar") + "") + ("a" + "foo bar")), "hello"}
    var v6 int16 = 1000
    var v17 int64 = 0
    for fi18 := 0; fi18 < 18; fi18++ {
        v5 = append(v5, ("\n\t" + ("foo bar" + (("hello" + ((((("a" + "") + "hello") + "") + "hello") + "hello")) + "\n\t"))))
        if (fi18 > fi18) {
            fmt.Println(v17, fi18, v17)
            var v19 bool = ((v4 % ((v4) | 1)) > (v4 ^ v4))
            _ = len(v5)
            _ = v19
        } else {
            var v20 uint = 41
            var v21 float32 = (0.0 / 0.1)
            _ = v20
            _ = v21
        }
        delete(v2, -122)
        _ = fi18
    }
    var v22 int64 = 1
    var v23 bool = (false && true)
    v23 = (4294967295 <= 255)
    if false {
        var v24 uint64 = 1
        _ = v24
    }
    v23 = false
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v25 int16 = v6
    v1[v4] = 4294967295
    for fi26 := 0; fi26 < 14; fi26++ {
        fmt.Println(v22, v6)
        v27 := v1[(v4 + (v4 ^ 0))]
        for fi28 := 0; fi28 < 18; fi28++ {
            fmt.Println(v23, fi28, fi28)
            switch ("foo bar" + "") {
            case "foo bar":
                v29 := GenS1{F0: (v4 | (^(v4)))}
                _ = v29
            case "":
                v30 := GenS1{F0: (v4 & (^(v4)))}
                _ = v30
            }
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v17, v17, v25)
    switch 32767 {
    case 42:
        var v31 float64 = 0.0
        _ = v31
    case 9223372036854775807:
        if false {
            v22 -= -695
            v32 := GenS1{F0: 2}
            _ = v32
        }
    default:
        var v33 float64 = 0.1
        _ = v33
    }
    var v34 int32 = -697
    v4 -= (^(v4))
    if v23 {
        v4 -= 638
        _ = len(v5)
        switch ("a" + (("\n\t" + "a") + "\n\t")) {
        case "\n\t":
            var v35 uint64 = 42
            _ = v35
        case "":
            v34 -= v34
        case "a":
            var v36 string = (((("" + "hello") + "foo bar") + "hello") + ("" + ("\n\t" + "hello")))
            _ = v36
        }
    }
    v23 = (32767 < -409)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
