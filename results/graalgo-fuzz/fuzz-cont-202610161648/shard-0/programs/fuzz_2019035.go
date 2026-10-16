package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 uint64, p2 uint) string {
    if true {
        p2 = p2
    }
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint = 568
        var v8 bool = false
        p1 = p1
        _ = fi6
        _ = v7
        _ = v8
    }
    if true {
        v9 := []int{}
        p1 = p1
        for fi10 := 0; fi10 < 18; fi10++ {
            for fi11 := 0; fi11 < 8; fi11++ {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                p0 = p0
                _ = fi11
            }
            var v12 int16 = 1
            fmt.Println(fi10, p2, p0)
            _ = fi10
            _ = v12
        }
        _ = v9
    } else {
        var v13 int8 = (p0 << (((p2 << (((p2 / (((p2 + (p2 >> ((p2) & 63)))) | 1))) & 63))) & 63))
        if (v13 < (v13 & v13)) {
            v14 := GenS1{F0: true}
            fmt.Println(p1, p2, p2)
            v15 := []int{4294967295, 9223372036854775807}
            _ = v14
            _ = v15
        } else {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: ((103 != 127) && ((!(true || true)) && false))}
                v18 := GenS1{F0: (2147483647 != 4294967295)}
                p0 = -12
                _ = fi16
                _ = v17
                _ = v18
            }
            v19 := []int{-1}
            v20 := GenS2{F0: 0}
            _ = v19
            _ = v20
        }
        p2 = 255
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return ("foo bar" + (("a" + ("foo bar" + "\n\t")) + "hello"))
}
func genFn1(p0 uint32) int64 {
    var v21 uint64 = 0
    if true {
        var v22 uint8 = 99
        for fi23 := 0; fi23 < 5; fi23++ {
            if false {
                v22 -= (v22 - v22)
                var v24 uint32 = p0
                v25 := GenS1{F0: false}
                _ = v24
                _ = v25
            } else {
                var v26 int8 = 80
                _ = v26
            }
            _ = fi23
        }
        for fi27 := 0; fi27 < 20; fi27++ {
            v22 -= v22
            v22 = (^(v22))
            _ = fi27
        }
        _ = v22
    }
    _ = p0
    _ = v21
    return 1
}

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 0
    var v3 uint64 = 42
    v4 := make(map[uint32]int)
    var v5 int16 = -1
    if ((!true) && (v5 > v5)) {
        var v28 uint64 = 2
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := GenS2{F0: -1}
            var v31 uint64 = v3
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = v28
    } else {
        v32 := GenS1{F0: true}
        _ = v32
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 uint8 = 85
        _ = fi33
        _ = v34
    }
    fmt.Println(v5, v2, v5)
    for fi35 := 0; fi35 < 18; fi35++ {
        switch (((("\n\t" + "foo bar") + "") + "") + ("hello" + (("foo bar" + "\n\t") + ""))) {
        case "a":
            switch 128 {
            case 2147483647:
                for fi36 := 0; fi36 < 8; fi36++ {
                    _ = len(v4)
                    _ = fi36
                }
            }
        default:
            var v37 uint16 = 176
            _ = v37
        }
        switch 128 {
        case 1:
            _ = len(v4)
        case -958:
            switch fi35 {
            case -357:
                fmt.Println(v3)
            case 4294967295:
                fmt.Println(v5, v1, fi35)
            default:
                fmt.Println(v1, fi35, v1)
            }
        }
        var v38 uint8 = 2
        _ = fi35
        _ = v38
    }
    v5 -= int16(v3)
    v39 := []int{}
    v3 -= 42
    for fi40 := 0; fi40 < 20; fi40++ {
        fmt.Println(v1, v3)
        _ = fi40
    }
    for fi41 := 0; fi41 < 12; fi41++ {
        v4[65535] = fi41
        _ = fi41
    }
    if (false && ((v3 == v3) && (-905 >= genFn1(2147483647)))) {
        v2 = (v2 % (((v2 ^ v2)) | 1))
    } else {
        v42 := GenS1{F0: true}
        _ = v42
    }
    var v43 uint32 = 318
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
