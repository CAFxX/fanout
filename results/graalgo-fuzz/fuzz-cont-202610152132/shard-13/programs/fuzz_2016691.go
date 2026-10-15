package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8, p1 bool) bool {
    var v4 string = "\n\t"
    if (!((!p1) || p1)) {
        for fi5 := 0; fi5 < 18; fi5++ {
            v6 := []int{(fi5 & fi5)}
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS1{F0: (p0 + 0), F1: 42, F2: (p0 >> ((255) & 63))}
                v9 := GenS1{F0: p0, F1: 4294967295, F2: (^(p0))}
                v10 := GenS1{F0: p0, F1: 2, F2: p0}
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi5
            _ = v6
        }
        switch v4 {
        case "foo bar":
            v11 := []int{-843}
            _ = v11
        case "":
            v12 := GenS1{F0: 42, F1: 4294967295, F2: (^(p0))}
            _ = v12
        case "hello":
            var v13 bool = p1
            _ = v13
        }
    } else {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint16 = 569
            _ = fi14
            _ = v15
        }
        var v16 uint = 65535
        if true {
            switch 0 {
            case -2147483648:
                fmt.Println(p1, p0, p0)
            }
        } else {
            var v17 int16 = 0
            _ = v17
        }
        _ = v16
    }
    p1 = false
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := []int{fi18, 909, fi18}
        _ = fi18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v4
    return false
}
func genFn1() uint32 {
    var v20 int = -122
    v20 = v20
    var v21 int32 = 1
    _ = v20
    _ = v21
    return 444
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]int)
    v3 := []string{"", "foo bar"}
    fmt.Println(v1)
    v22 := GenS1{F0: 127, F1: (v1 - 18446744073709551615), F2: 203}
    var v23 bool = true
    switch ("\n\t" + "foo bar") {
    case "hello":
        fmt.Println(v23, v1)
    case "":
        fmt.Println(v23)
    }
    if false {
        var v24 int = 9223372036854775807
        var v25 int32 = 1
        _ = v24
        _ = v25
    }
    v23 = (false || (1 < -1))
    var v26 int = 416
    if (32767 != 32767) {
        if (((42 >= -9223372036854775808) || ((v26 ^ (v26 * -455)) == (v26 - v26))) || ((584 <= 255) || false)) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v27 := v2[genFn1()]
        _ = v27
    }
    if false {
        v28 := GenS2{F0: ((0 == -32768) && true)}
        var v29 int32 = 42
        _ = v28
        _ = v29
    }
    v26 = (v26 << ((161) & 63))
    var v30 uint16 = 816
    v31 := GenS2{F0: true}
    v1 = 18446744073709551615
    v32 := GenS2{F0: (1e10 > (1.5 + (1e10 * 0.1)))}
    for fi33 := 0; fi33 < 2; fi33++ {
        v30 += 42
        var v34 int8 = 42
        _ = fi33
        _ = v34
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if v23 {
        if (false && ((v26 ^ (v26 * (^(v26)))) > (v26 % ((v26) | 1)))) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v26 = 32768
            if (217 >= 1) {
                fmt.Println(v23, v23, v1)
                delete(v2, 1)
                var v35 uint64 = (v1 + (^(v1)))
                _ = v35
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
