package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint16, p2 int32) int8 {
    p0 = p0
    var v7 uint32 = 65535
    p2 -= -2147483648
    v8 := GenS1{F0: true}
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]bool)
    var v3 uint = 693
    var v4 uint8 = 128
    var v5 int64 = 1
    var v6 float64 = (0.0 - (-2.25 * 0.0))
    switch 128 {
    case 42:
        if true {
            if false {
                if true {
                    v10 := GenS1{F0: false}
                    fmt.Println(v4)
                    _ = v10
                }
                for fi11 := 0; fi11 < 4; fi11++ {
                    for fi12 := 0; fi12 < 16; fi12++ {
                        fmt.Println(fi12)
                        v13 := GenS1{F0: true}
                        _ = fi12
                        _ = v13
                    }
                    delete(v2, 2147483647)
                    _ = fi11
                }
            }
            var v14 float64 = v6
            _ = v14
        }
    case 127:
        fmt.Println(v3)
    case -574:
        fmt.Println(v4)
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        if (true && true) {
            fmt.Println(v1)
            for fi16 := 0; fi16 < 1; fi16++ {
                switch (fi15 - (fi16 * fi16)) {
                case -1:
                    switch ("a" + ("foo bar" + ("\n\t" + "hello"))) {
                    case "a":
                        v17 := []int{}
                        _ = v17
                    case "hello":
                        v5 -= (v5 ^ v5)
                    default:
                        v18 := GenS2{F0: 397}
                        _ = v18
                    }
                case 4294967295:
                    v6 = v6
                }
                if (true || (0 == 2)) {
                    fmt.Println(v6, v3)
                }
                _ = fi16
            }
            v5 = (v5 & (v5 + (^(v5))))
        } else {
            v1 = -1
            var v19 string = ("hello" + "\n\t")
            _ = v19
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := []int{}
            v22 := v2[42]
            v1 = (v1 & (v1 - 792))
            _ = fi20
            _ = v21
            _ = v22
        }
        if false {
            if (117 == 2147483648) {
                var v23 uint = v3
                var v24 int16 = 0
                v1 += (v24 ^ (v1 & v1))
                _ = v23
                _ = v24
            } else {
                var v25 string = "\n\t"
                if true {
                    switch (fi15 / ((fi15) | 1)) {
                    case 42:
                        fmt.Println(fi15, v1, v3)
                    case -413:
                        for fi26 := 0; fi26 < 18; fi26++ {
                            var v27 int = fi26
                            fmt.Println(fi15, v4, v1)
                            _ = fi26
                            _ = v27
                        }
                    case -248:
                        switch (fi15 & fi15) {
                        case 4294967295:
                            v28 := []int{(^(fi15)), 42, fi15, (fi15 - fi15)}
                            _ = v28
                        case 128:
                            v29 := []int{fi15, fi15, (^(fi15))}
                            _ = v29
                        }
                    default:
                        var v30 uint8 = (^(v4))
                        _ = v30
                    }
                }
                v31 := v2[2147483647]
                _ = v25
                _ = v31
            }
            var v32 int8 = 0
            _ = v32
        }
        _ = fi15
    }
    var v33 uint = v3
    v34 := []int{}
    v2[159] = false
    v1 = 1
    v35 := GenS2{F0: 449}
    v36 := v2[0]
    if true {
        if (3.14159 == (0.1 * ((float64(-2.25) + v6) - 1e10))) {
            v37 := v2[4294967295]
            fmt.Println(v37, v37)
            _ = v37
        } else {
            var v38 uint8 = v4
            v39 := v2[2147483648]
            _ = v38
            _ = v39
        }
        v36 = v36
    }
    v3 += v33
    if ((9223372036854775807 == -439) || v36) {
        var v40 float32 = (0.0 / 3.14159)
        _ = v40
    }
    v41 := GenS1{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
