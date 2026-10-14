package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 string }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 int64, p1 string) uint16 {
    for fi4 := 0; fi4 < 20; fi4++ {
        if (!true) {
            if true {
                fmt.Println(fi4, fi4)
                var v5 uint32 = 65535
                _ = v5
            } else {
                v6 := []int{fi4, fi4, -362}
                _ = v6
            }
            p1 = p1
            if false {
                v7 := []int{(^(fi4)), fi4}
                p1 = ((((((p1 + p1) + "") + p1) + ("\n\t" + "hello")) + p1) + "a")
                var v8 float64 = float64((1e10 - 1.5))
                _ = v7
                _ = v8
            } else {
                v9 := GenS1{F0: 331, F1: 2, F2: "foo bar"}
                for fi10 := 0; fi10 < 11; fi10++ {
                    fmt.Println(p0, fi4, fi4)
                    _ = fi10
                }
                _ = v9
            }
        } else {
            p1 = ((p1 + ((p1 + "a") + p1)) + p1)
            v11 := []int{32768, fi4, fi4}
            var v12 uint = 975
            _ = v11
            _ = v12
        }
        _ = fi4
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 397
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ("foo bar" + "hello")
    var v3 int = -1
    v13 := GenS1{F0: 0, F1: 0, F2: (v2 + ((v2 + v2) + ("" + v2)))}
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := v1[(fi14 >> ((255) & 63))]
        _ = fi14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                v2 = "a"
                v18 := GenS1{F0: 184, F1: 255, F2: v2}
                v1[(fi16 + fi17)] = 0
                _ = fi17
                _ = v18
            }
            var v19 uint16 = 255
            v3 = 32768
            _ = fi16
            _ = v19
        }
    }
    if true {
        var v20 int64 = 1
        _ = v20
    }
    delete(v1, v3)
    var v21 float64 = float64((1e10 - (-2.25 - 1.5)))
    v22 := GenS2{F0: 32767, F1: v2}
    var v23 int = 128
    switch (v23 * v3) {
    case 127:
        v3 = (-(v23))
    case 32768:
        if true {
            if false {
                var v24 int = v23
                _ = v24
            }
        }
    case -1:
        v3 -= (v23 * (v23 << ((42) & 63)))
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        v2 = v2
        _ = fi25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
