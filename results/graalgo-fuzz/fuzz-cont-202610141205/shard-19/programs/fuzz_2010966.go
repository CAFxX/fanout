package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
func genFn0(p0 bool, p1 uint16, p2 int64) int16 {
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint = 4294967295
        switch (fi5 ^ -9223372036854775808) {
        case -2147483648:
            for fi7 := 0; fi7 < 5; fi7++ {
                if ((255 >= 181) || p0) {
                    var v8 int16 = 345
                    v6 = (v6 & 645)
                    var v9 float64 = -2.25
                    _ = v8
                    _ = v9
                }
                p1 -= 1
                fmt.Println(v6, p1)
                _ = fi7
            }
        }
        _ = fi5
        _ = v6
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            p2 -= p2
            if (77 >= 127) {
                p2 -= p2
                if (1 >= 42) {
                    fmt.Println(p0, fi10, fi11)
                }
            }
            switch fi10 {
            case -348:
                var v12 uint16 = 32767
                _ = v12
            case 434:
                p0 = false
            case 42:
                v13 := []int{fi10, (fi11 << ((65535) & 63)), 4294967295, fi11}
                var v14 string = ((("hello" + "foo bar") + "foo bar") + "hello")
                switch fi11 {
                case -32:
                    fmt.Println(fi10, v14, v14)
                }
                _ = v13
                _ = v14
            }
            _ = fi11
        }
        _ = fi10
    }
    v15 := []int{32768}
    var v16 int32 = 906
    _ = len(v15)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 1000
}
func genFn1(p0 int, p1 int64) uint8 {
    var v17 uint8 = 1
    var v18 int16 = -149
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return (v17 % ((13) | 1))
}

func main() {
    var v1 int32 = -1
    var v2 string = (((((((("\n\t" + ("\n\t" + "\n\t")) + "foo bar") + "hello") + ("hello" + "foo bar")) + "") + (("\n\t" + "a") + "")) + "foo bar") + "hello")
    var v3 uint = 0
    v4 := []string{}
    var v19 int32 = v1
    v1 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 int16 = 280
        var v22 uint32 = 2147483647
        _ = fi20
        _ = v21
        _ = v22
    }
    var v23 uint32 = 94
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 uint32 = (^(v23))
        _ = fi24
        _ = v25
    }
    var v26 uint64 = 450
    for fi27 := 0; fi27 < 6; fi27++ {
        v1 = v19
        v28 := GenS1{F0: -9223372036854775808, F1: 661}
        _ = fi27
        _ = v28
    }
    var v29 int16 = 0
    var v30 uint32 = v23
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
