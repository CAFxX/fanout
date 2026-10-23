package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 bool) int {
    v6 := []int{}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 string = "\n\t"
        if len(v6) > 0 {
            v6[0] = fi7
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = v6
    return 32768
}

func main() {
    var v1 uint = 721
    v2 := make(map[int]int)
    var v3 uint32 = 2147483648
    v4 := make(map[uint32]bool)
    var v5 string = "foo bar"
    fmt.Println(v1)
    fmt.Println(v5)
    fmt.Println(v1)
    var v9 uint = 42
    var v10 int32 = 42
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint8 = 152
        var v13 int16 = 1000
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint = v1
    var v15 int64 = 1
    switch (v5 + (v5 + "a")) {
    case "":
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 13; fi17++ {
                if (true || true) {
                    v18 := v2[(^(fi17))]
                    _ = v18
                } else {
                    v19 := []int{(fi16 / ((genFn0(((!(!((^(v15)) <= (v15 ^ v15)))) || true))) | 1)), fi16, fi17}
                    _ = v19
                }
                _ = fi17
            }
            if true {
                v20 := GenS1{F0: 18446744073709551615}
                for fi21 := 0; fi21 < 13; fi21++ {
                    v22 := GenS1{F0: 4294967295}
                    _ = fi21
                    _ = v22
                }
                v3 = 2147483648
                _ = v20
            } else {
                v23 := GenS1{F0: 4294967295}
                v24 := []int{fi16, 457, genFn0(false)}
                v10 = v10
                _ = v23
                _ = v24
            }
            _ = fi16
        }
    case "hello":
        v1 += v14
    }
    v14 = (v1 >> ((v9) & 63))
    for fi25 := 0; fi25 < 4; fi25++ {
        fmt.Println(v1)
        _ = fi25
    }
    var v26 uint32 = v3
    if true {
        switch 42 {
        case 36:
            if (((!false) && (!(4294967295 < 4294967295))) && (("" < ("foo bar" + v5)) && true)) {
                v27 := v4[42]
                for fi28 := 0; fi28 < 12; fi28++ {
                    var v29 int = (fi28 / ((32767) | 1))
                    v1 -= v14
                    _ = fi28
                    _ = v29
                }
                _ = v27
            }
        case 4294967295:
            v30 := GenS2{F0: 1}
            _ = v30
        case 444:
            var v31 uint32 = (^(v26))
            _ = v31
        }
        fmt.Println(v10, v9, v3)
        var v32 string = v5
        _ = v32
    } else {
        for fi33 := 0; fi33 < 16; fi33++ {
            switch ("foo bar" + v5) {
            case "a":
                var v34 uint = v14
                _ = v34
            case "foo bar":
                v9 = v9
            case "hello":
                var v35 float32 = 1e10
                _ = v35
            default:
                var v36 int32 = v10
                _ = v36
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                var v38 string = v5
                _ = fi37
                _ = v38
            }
            v39 := v2[-994]
            _ = fi33
            _ = v39
        }
        v5 = v5
        var v40 bool = ((988 == 0) || (!false))
        _ = v40
    }
    v41 := []int{9223372036854775807, genFn0(((v9 | v1) == v14)), 128, 32767}
    var v42 int = 577
    for fi43 := 0; fi43 < 17; fi43++ {
        var v44 uint = 255
        if true {
            var v45 bool = false
            _ = v45
        }
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v9
}
